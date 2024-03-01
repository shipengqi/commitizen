package render

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shipengqi/commitizen/internal/ui"
)

type Option struct {
	Name string
	Desc string
}

func (o *Option) String() string {
	var b strings.Builder
	ml := len(o.Name)
	pl := 10 - ml - 2
	padding := strings.Repeat(" ", pl)
	b.WriteString(o.Name)
	b.WriteString(": ")
	b.WriteString(padding)
	b.WriteString(o.Desc)
	return b.String()
}

type Item struct {
	Name     string
	Desc     string
	Type     string
	Options  []Option
	Required bool
}

type Template struct {
	Name   string
	Format string
	Items  []*Item
	models []model
}

type model struct {
	t     string
	name  string
	model ui.Model
}

func NewTemplate() (*Template, error) {
	t := &Template{}
	err := t.init()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Template) Run() ([]byte, error) {
	err := t.init()
	if err != nil {
		return nil, err
	}

	if len(t.models) == 0 {
		return nil, nil
	}

	values := map[string]interface{}{}
	for _, v := range t.models {
		if _, err = tea.NewProgram(v.model).Run(); err != nil {
			return nil, err
		}
		if v.model.Canceled() {
			return nil, errors.New("canceled")
		}
		val := v.model.Value()
		// hardcode for the select options
		if v.t == TypeSelect {
			tokens := strings.Split(val, ":")
			if len(tokens) > 0 {
				val = tokens[0]
			}
		}
		values[v.name] = val
	}

	tmpl, err := template.New("cz").Parse(t.Format)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, values); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *Template) init() error {
	if isEmptyStr(t.Format) {
		return errors.New("format is required")
	}

	for _, item := range t.Items {
		if isEmptyStr(item.Name) {
			return errors.New("item.name is required")
		}
		if isEmptyStr(item.Desc) {
			return errors.New("item.desc is required")
		}
		if isEmptyStr(item.Type) {
			return errors.New("item.type is required")
		}

		var m ui.Model

		switch item.Type {
		case TypeInput:
			m = t.createInputItem(item.Name, item.Desc, item.Required)
		case TypeSelect:
			m = t.createSelectItem(item.Desc, item.Options)
		case TypeTextArea:
			m = t.createTextAreaItem(item.Name, item.Desc, item.Required)
		default:
			return fmt.Errorf("unsupported type: %s", item.Type)
		}
		t.models = append(t.models, model{
			t:     item.Type,
			name:  item.Name,
			model: m,
		})
	}
	return nil
}

func (t *Template) createSelectItem(label string, options []Option) *ui.SelectModel {
	var choices ui.Choices
	for _, v := range options {
		choices = append(choices, ui.Choice(v.String()))
	}
	m := ui.NewSelect(label, choices)
	return m
}

func (t *Template) createInputItem(name, label string, required bool) *ui.InputModel {
	m := ui.NewInput(label)
	if required {
		m.WithValidateFunc(NotBlankValidator(name))
	}
	return m
}

func (t *Template) createTextAreaItem(name, label string, required bool) *ui.TextAreaModel {
	m := ui.NewTextArea(label)
	if required {
		m.WithValidateFunc(NotBlankValidator(name))
	}
	return m
}

// NotBlankValidator is a verification function that checks whether the input is empty
func NotBlankValidator(name string) func(s string) error {
	return func(s string) error {
		if strings.TrimSpace(s) == "" {
			return fmt.Errorf("%s is required", name)
		}
		return nil
	}
}
