package list

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/errors"
	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/validators"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	Options      []huh.Option[string] `yaml:"options"       json:"options"       mapstructure:"options"`
	DefaultValue string               `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
	Required     bool                 `yaml:"required"      json:"required"      mapstructure:"required"`
}

func (p Param) Validate() []error {
	errs := p.Parameter.Validate()
	if len(p.Options) < 1 {
		errs = append(errs, errors.NewMissingErr("options", p.Name))
	}
	return errs
}

func (p Param) Render() huh.Field {
	param := huh.NewSelect[string]().Key(p.Name).
		Options(p.Options...).
		Title(p.Label)
	if len(p.Description) > 0 {
		param.Description(p.Description)
	}

	param.Value(&p.DefaultValue)

	var group []validators.Validator
	if p.Required {
		group = append(group, validators.Required(p.Name, false))
	}

	if len(group) > 0 {
		param.Validate(validators.Group(group...))
	}
	return param
}
