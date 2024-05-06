package parameter

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/shipengqi/golib/strutil"

	"github.com/shipengqi/commitizen/internal/errors"
)

type Interface interface {
	huh.Field
	GetGroup() string
	Render()
	Validate() []error
}

type Parameter struct {
	huh.Field   `mapstructure:"-"`
	Name        string    `yaml:"name"        json:"name"        mapstructure:"name"`
	Group       string    `yaml:"group"       json:"group"       mapstructure:"group"`
	Label       string    `yaml:"label"       json:"label"       mapstructure:"label"`
	Description string    `yaml:"description" json:"description" mapstructure:"description"`
	Type        string    `yaml:"type"        json:"type"        mapstructure:"type"`
	DependsOn   DependsOn `yaml:"depends_on"  json:"depends_on"  mapstructure:"depends_on"`
}

func (p *Parameter) GetGroup() string {
	return p.Group
}

func (p *Parameter) Render() {}

func (p *Parameter) Validate() []error {
	var errs []error
	if strutil.IsEmpty(p.Name) {
		errs = append(errs, errors.NewMissingErr("parameter.name"))
	}
	if !regexName.MatchString(p.Name) {
		errs = append(errs, fmt.Errorf("parameter.name '%s' must match the regex: ^[a-zA-Z0-9-_]{1,62}$", p.Name))
	}
	if strutil.IsEmpty(p.Label) {
		errs = append(errs, errors.NewMissingErr("label", p.Name))
	}
	if strutil.IsEmpty(p.Type) {
		errs = append(errs, errors.NewMissingErr("type", p.Name))
	}
	return errs
}
