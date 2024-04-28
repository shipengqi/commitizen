package str

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/validators"
)

type Param struct {
	parameter.Parameter `mapstructure:",squash"`

	Required     bool   `yaml:"required"      json:"required"      mapstructure:"required"`
	FQDN         bool   `yaml:"fqdn"          json:"fqdn"          mapstructure:"fqdn"`
	IP           bool   `yaml:"ip"            json:"ip"            mapstructure:"ip"`
	Trim         bool   `yaml:"trim"          json:"trim"          mapstructure:"trim"`
	DefaultValue string `yaml:"default_value" json:"default_value" mapstructure:"default_value"`
	Regex        string `yaml:"regex"         json:"regex"         mapstructure:"regex"`
	MinLength    *int   `yaml:"min_length"    json:"min_length"    mapstructure:"min_length"`
	MaxLength    *int   `yaml:"max_length"    json:"max_length"    mapstructure:"max_length"`
}

func (p *Param) Validate() []error {
	errs := p.Parameter.Validate()
	if p.Regex != "" {
		if _, err := regexp.Compile(p.Regex); err != nil {
			errs = append(errs, fmt.Errorf("regex %s compile: %s", p.Regex, err.Error()))
		}
	}
	return errs
}

func (p *Param) Render() {
	p.Field = p.RenderInput()
}

func (p *Param) GetValue() any {
	if !p.Trim {
		return p.Field.GetValue()
	}
	val := p.Field.GetValue()
	if str, ok := val.(string); ok {
		return strings.TrimSpace(str)
	}
	return p.Field.GetValue()
}

func (p *Param) RenderInput() *huh.Input {
	param := huh.NewInput().Key(p.Name).
		Title(p.Label)

	if len(p.Description) > 0 {
		param.Description(p.Description)
	}

	param.Value(&p.DefaultValue)

	var group []validators.Validator[string]
	if p.Required {
		group = append(group, validators.Required(p.Name, p.Trim))
	}
	// if the value is not required and no value has been given, min length validator should be ignored.
	if p.Required && p.MinLength != nil {
		group = append(group, validators.MinLength(*p.MinLength))
	}
	if p.MaxLength != nil {
		group = append(group, validators.MaxLength(*p.MaxLength))
	}
	if p.IP {
		group = append(group, validators.IPValidator())
	}
	if p.FQDN {
		group = append(group, validators.FQDNValidator())
	}
	if p.Regex != "" {
		group = append(group, validators.RegexValidator(p.Regex))
	}

	if len(group) > 0 {
		param.Validate(validators.Group(group...))
	}
	return param
}
