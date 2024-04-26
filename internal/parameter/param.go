package parameter

import (
	"github.com/charmbracelet/huh"
	"github.com/shipengqi/golib/strutil"

	"github.com/shipengqi/commitizen/internal/errors"
)

type Interface interface {
	GetGroup() string
	Render() huh.Field
	Validate() []error
}

type Parameter struct {
	Name        string `yaml:"name"        json:"name"        mapstructure:"name"`
	Group       string `yaml:"group"       json:"group"       mapstructure:"group"`
	Label       string `yaml:"label"       json:"label"       mapstructure:"label"`
	Description string `yaml:"description" json:"description" mapstructure:"description"`
	Type        string `yaml:"type"        json:"type"        mapstructure:"type"`
	// DependsOn   DependsOn `yaml:"depends_on"  json:"depends_on"  mapstructure:"depends_on"`
}

func (p Parameter) GetGroup() string {
	return p.Group
}

func (p Parameter) Render() huh.Field {
	return nil
}

func (p Parameter) Validate() []error {
	var errs []error
	if strutil.IsEmpty(p.Name) {
		errs = append(errs, errors.NewRequiredErr("parameter.name"))
	}
	if strutil.IsEmpty(p.Label) {
		errs = append(errs, errors.NewRequiredErr("parameter.label"))
	}
	if strutil.IsEmpty(p.Type) {
		errs = append(errs, errors.NewRequiredErr("parameter.type"))
	}
	return errs
}

type DependsOn struct {
	AndConditions []Condition `yaml:"and_conditions"  json:"and_conditions" mapstructure:"and_conditions"`
	OrConditions  []Condition `yaml:"or_conditions"   json:"or_conditions"  mapstructure:"or_conditions"`
}

type Condition interface {
	Match() bool
}

type EqualsCondition struct {
	ParameterName string `yaml:"parameter_name" json:"parameter_name" mapstructure:"parameter_name"`
	ValueEquals   string `yaml:"value_equals"   json:"value_equals"   mapstructure:"value_equals"`
}

type NotEqualsCondition struct {
	ParameterName string `yaml:"parameter_name" json:"parameter_name" mapstructure:"parameter_name"`
	ValueContains string `yaml:"value_contains" json:"value_contains" mapstructure:"value_contains"`
}

type ContainsCondition struct {
	ParameterName    string `yaml:"parameter_name"     json:"parameter_name"     mapstructure:"parameter_name"`
	ValueNotContains string `yaml:"value_not_contains" json:"value_not_contains" mapstructure:"value_not_contains"`
}

type NotContainsCondition struct {
	ParameterName string `yaml:"parameter_name" json:"parameter_name" mapstructure:"parameter_name"`
	ValueEmpty    string `yaml:"value_empty"    json:"value_empty"    mapstructure:"value_empty"`
}

type EmptyCondition struct {
	ParameterName string `yaml:"parameter_name" json:"parameter_name" mapstructure:"parameter_name"`
	ValueEquals   string `yaml:"value_equals"   json:"value_equals"   mapstructure:"value_equals"`
}
