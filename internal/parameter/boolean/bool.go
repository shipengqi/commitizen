package boolean

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	DefaultValue bool `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
}

func (p *Param) Render() {
	param := huh.NewConfirm().Key(p.Name).
		Title(p.Label)

	if len(p.Description) > 0 {
		param.Description(p.Description)
	}

	param.Value(&p.DefaultValue)

	p.Field = param
}
