package config

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shipengqi/golib/convutil"
	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/golib/sysutil"
	"gopkg.in/yaml.v3"

	"github.com/shipengqi/commitizen/internal/render"
	"github.com/shipengqi/commitizen/internal/ui"
)

const RCFilename = ".git-czrc"

type Config struct {
	defaultTmpl *render.Template
	others      []*render.Template
}

func New() *Config {
	return &Config{}
}

func (c *Config) Initialize() error {
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
		return err
	}
	for _, v := range tmpls {
		if v.Default {
			if c.defaultTmpl != nil {
				// the default template already exists
				return errors.New("only one default template is permitted")
			}
			c.defaultTmpl = v
			continue
		}
		c.others = append(c.others, v)
	}
	// If the user has not configured a default template, the built-in template is used
	if c.defaultTmpl == nil {
		defaults, err := Load(convutil.S2B(DefaultCommitTemplate))
		if err != nil {
			return err
		}
		c.defaultTmpl = defaults[0]
	}

	return nil
}

func (c *Config) Run() (*render.Template, error) {
	if len(c.others) > 0 {
		model := c.createTemplatesSelect("Select the template to be used for this commit")
		if _, err := tea.NewProgram(model).Run(); err != nil {
			return nil, err
		}
		if model.Canceled() {
			return nil, render.ErrCanceled
		}
		val := model.Value()
		if val == c.defaultTmpl.Name {
			return c.defaultTmpl, nil
		}
		for _, v := range c.others {
			if v.Name == val {
				return v, nil
			}
		}
	}
	return c.defaultTmpl, nil
}

func (c *Config) createTemplatesSelect(label string) *ui.SelectModel {
	var choices ui.Choices
	var all []*render.Template
	all = append(all, c.others...)
	all = append(all, c.defaultTmpl)
	// list custom templates and default templates
	for _, v := range all {
		choices = append(choices, ui.Choice(v.Name))
	}
	height := 8
	if len(all) > 5 {
		height = 12
	} else if len(all) > 3 {
		height = 10
	} else if len(all) > 2 {
		height = 9
	}
	m := ui.NewSelect(label, choices).WithHeight(height)
	return m
}

func LoadTemplates(file string) ([]*render.Template, error) {
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

func Load(data []byte) ([]*render.Template, error) {
	return load(bytes.NewReader(data))
}

func load(reader io.Reader) ([]*render.Template, error) {
	var templates []*render.Template
	d := yaml.NewDecoder(reader)
	for {
		tmpl := new(render.Template)
		err := d.Decode(tmpl)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		templates = append(templates, tmpl)
	}

	return templates, nil
}
