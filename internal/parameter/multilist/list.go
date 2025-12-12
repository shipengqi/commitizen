package multilist

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/errorsx"
	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/validators"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	Options      []huh.Option[string] `yaml:"options"       json:"options"       mapstructure:"options"`
	DefaultValue []string             `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
	Required     bool                 `yaml:"required"      json:"required"      mapstructure:"required"`
	Limit        *int                 `yaml:"limit"         json:"limit"         mapstructure:"limit"`
	Height       *int                 `yaml:"height"        json:"height"        mapstructure:"height"`
}

func (p *Param) Validate() []error {
	errs := p.Parameter.Validate()
	if len(p.Options) < 1 {
		errs = append(errs, errorsx.NewMissingErr("options", p.Name))
	}
	return errs
}

func (p *Param) Render() {
	param := huh.NewMultiSelect[string]().Key(p.Name).
		Options(p.Options...).
		Title(p.Label)
	if len(p.Description) > 0 {
		param.Description(p.Description)
	}
	if p.Height != nil {
		param.Height(*p.Height)
	}
	if p.Limit != nil {
		param.Limit(*p.Limit)
	}
	param.Value(&p.DefaultValue)

	var group []validators.Validator[[]string]
	if p.Required {
		group = append(group, validators.MultiRequired(p.Name))
	}

	if p.Required {
		param.Validate(validators.Group(group...))
	}

	p.Field = param
}
