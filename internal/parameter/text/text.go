package text

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/validators"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	Required     bool   `yaml:"required"      json:"required"      mapstructure:"required"`
	DefaultValue string `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
	Regex        string `yaml:"regex"         json:"regex"         mapstructure:"regex"`
	RegexMessage string `yaml:"regex_message" json:"regex_message" mapstructure:"regex_message"`
	MinLength    *int   `yaml:"min_length"    json:"min_length"    mapstructure:"min_length"`
	MaxLength    *int   `yaml:"max_length"    json:"max_length"    mapstructure:"max_length"`
	Height       *int   `yaml:"height"        json:"height"        mapstructure:"height"`
}

func (p Param) Render() huh.Field {
	param := huh.NewText().Key(p.Name).
		Title(p.Label)

	if p.Height != nil {
		param.Lines(*p.Height)
	}

	if len(p.Description) > 0 {
		param.Description(p.Description)
	}

	param.Value(&p.DefaultValue)

	var group []validators.Validator
	if p.Required {
		group = append(group, validators.Required(p.Name, false))
	}
	if p.MinLength != nil {
		group = append(group, validators.MinLength(*p.MinLength))
	}
	if p.MaxLength != nil {
		group = append(group, validators.MaxLength(*p.MaxLength))
	}
	if p.Regex != "" {
		group = append(group, validators.RegexValidator(p.Regex, p.RegexMessage))
	}

	if len(group) > 0 {
		param.Validate(validators.Group(group...))
	}

	return param
}
