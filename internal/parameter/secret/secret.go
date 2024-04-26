package secret

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter/str"
)

type Param struct {
	str.Param `mapstructure:",squash"`
}

func (p Param) Render() huh.Field {
	input := p.Param.RenderInput()
	input.Password(true)
	return input
}
