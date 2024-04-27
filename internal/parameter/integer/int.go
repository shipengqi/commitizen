package integer

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/validators"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	DefaultValue string `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
	Required     bool   `yaml:"required"      json:"required"      mapstructure:"required"`
	Min          *int   `yaml:"min"           json:"min"           mapstructure:"min"`
	Max          *int   `yaml:"max"           json:"max"           mapstructure:"max"`
}

func (p Param) Render() huh.Field {
	param := huh.NewInput().Key(p.Name).
		Title(p.Label)

	if len(p.Description) > 0 {
		param.Description(p.Description)
	}

	param.Value(&p.DefaultValue)

	var group []validators.Validator[string]
	if p.Required {
		group = append(group, validators.Required(p.Name, true))
	}
	if p.Min != nil {
		group = append(group, validators.Min(*p.Min))
	}
	if p.Max != nil {
		group = append(group, validators.Max(*p.Max))
	}

	if len(group) > 0 {
		param.Validate(validators.Group(group...))
	}
	return param
}
