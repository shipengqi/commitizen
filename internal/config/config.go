package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/shipengqi/golib/convutil"
	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/golib/sysutil"
	"gopkg.in/yaml.v3"

	"github.com/shipengqi/commitizen/internal/options"
	"github.com/shipengqi/commitizen/internal/templates"
)

const (
	RCFilename             = ".ggczrc"
	ReservedDefaultName    = "default"
	FieldKeyTemplateSelect = "template-select"
)

type Config struct {
	defaultTmpl *templates.Template
	more        []*templates.Template
}

func New() *Config {
	return &Config{}
}

func (c *Config) initialize() error {
	var fpath string
	if fsutil.IsExists(RCFilename) {
		fpath = RCFilename
	} else {
		home := sysutil.HomeDir()
		p := filepath.Join(home, RCFilename)
		if fsutil.IsExists(p) {
			fpath = p
		}
	}

	tmpls, err := LoadTemplates(fpath)
	if err != nil {
		return fmt.Errorf("load templates %s failed: %v", fpath, err)
	}
	exists := make(map[string]struct{}, len(tmpls))
	for _, v := range tmpls {
		if v.Default {
			if c.defaultTmpl != nil {
				// the default template already exists
				return errors.New("only one default template is permitted")
			}
			c.defaultTmpl = v
			continue
		}
		if v.Name == ReservedDefaultName {
			return errors.New("template name 'default' is reserved, to override the default template, you need to set default to true")
		}
		if _, ok := exists[v.Name]; ok {
			return fmt.Errorf("duplicate template '%s'", v.Name)
		}

		exists[v.Name] = struct{}{}
		c.more = append(c.more, v)
	}
	// If the user has not configured a default template, use the built-in template as the default template
	if c.defaultTmpl == nil {
		defaults, err := Load(convutil.S2B(DefaultCommitTemplate))
		if err != nil {
			return err
		}
		c.defaultTmpl = defaults[0]
	}

	return nil
}

func (c *Config) Run(opts *options.Options) (*templates.Template, error) {
	err := c.initialize()
	if err != nil {
		return nil, err
	}

	if opts.Default {
		return c.defaultTmpl, nil
	}
	// find the given template
	if len(opts.Template) > 0 {
		if opts.Template == c.defaultTmpl.Name {
			return c.defaultTmpl, nil
		}
		for _, v := range c.more {
			if v.Name == opts.Template {
				return v, nil
			}
		}
		return nil, fmt.Errorf("template '%s' not found", opts.Template)
	}

	if len(c.more) > 0 {
		form := c.createTemplatesSelect("Select the template of change that you're committing:")
		if err = form.Run(); err != nil {
			return nil, err
		}

		val := form.GetString(FieldKeyTemplateSelect)
		if val == c.defaultTmpl.Name {
			return c.defaultTmpl, nil
		}
		for _, v := range c.more {
			if v.Name == val {
				return v, nil
			}
		}
	}
	return c.defaultTmpl, nil
}

func (c *Config) createTemplatesSelect(label string) *huh.Form {
	var choices []string
	var all []*templates.Template
	all = append(all, c.more...)
	all = append(all, c.defaultTmpl)
	// list custom templates and default templates
	for _, v := range all {
		choices = append(choices, v.Name)
	}

	return huh.NewForm(huh.NewGroup(
		huh.NewNote().Title("Commitizen").Description("Welcome to Commitizen!\nFor further configuration visit:\nhttps://github.com/shipengqi/commitizen"),
		huh.NewSelect[string]().
			Key(FieldKeyTemplateSelect).
			Options(huh.NewOptions(choices...)...).
			Title(label)),
	)
}

func LoadTemplates(file string) ([]*templates.Template, error) {
	if len(file) == 0 {
		return nil, nil
	}
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() { _ = fd.Close() }()
	return load(fd)
}

func Load(data []byte) ([]*templates.Template, error) {
	return load(bytes.NewReader(data))
}

func load(reader io.Reader) ([]*templates.Template, error) {
	var tmpls []*templates.Template
	d := yaml.NewDecoder(reader)
	for {
		tmpl := new(templates.Template)
		err := d.Decode(tmpl)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		err = tmpl.Initialize()
		if err != nil {
			return nil, err
		}
		tmpls = append(tmpls, tmpl)
	}

	return tmpls, nil
}
