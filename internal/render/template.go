package render

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/charmbracelet/huh"
)

type Option struct {
	Name string
	Desc string
}

func (o *Option) String() string {
	var b strings.Builder
	ml := len(o.Name)
	pl := 12 - ml - 2
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
	Name    string
	Desc    string
	Format  string
	Default bool
	Items   []*Item
	fields  []huh.Field
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

	if len(t.fields) == 0 {
		return nil, nil
	}

	values := map[string]interface{}{}
	groups := ArrayInGroupsOf(t.fields, 2)
	form := huh.NewForm(groups...)
	err = form.Run()
	if err != nil {
		return nil, err
	}
	for _, field := range t.fields {
		values[field.GetKey()] = field.GetValue()
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
		return NewMissingErr("format")
	}

	exists := make(map[string]struct{}, len(t.Items))
	for _, item := range t.Items {
		if isEmptyStr(item.Name) {
			return NewMissingErr("item.name")
		}
		if isEmptyStr(item.Desc) {
			return NewMissingErr("item.desc")
		}
		if isEmptyStr(item.Type) {
			return NewMissingErr("item.type")
		}
		if _, ok := exists[item.Name]; ok {
			return fmt.Errorf("duplicate item.name: %s", item.Name)
		}
		exists[item.Name] = struct{}{}

		var field huh.Field

		switch item.Type {
		case TypeInput:
			field = t.createInputItem(item.Name, item.Desc, item.Required)
		case TypeSelect:
			field = t.createSelectItem(item.Name, item.Desc, item.Options)
		case TypeTextArea:
			field = t.createTextAreaItem(item.Name, item.Desc, item.Required)
		default:
			return fmt.Errorf("unsupported type: %s", item.Type)
		}
		t.fields = append(t.fields, field)
	}
	return nil
}

func (t *Template) createSelectItem(name, label string, options []Option, desc ...string) *huh.Select[string] {
	var choices []huh.Option[string]
	for _, v := range options {
		choices = append(choices, huh.Option[string]{
			Key:   v.String(),
			Value: v.Name,
		})
	}

	se := huh.NewSelect[string]().
		Key(name).
		Options(choices...).
		Title(label)
	if len(desc) > 0 && desc[0] != "" {
		se.Description(desc[0])
	}
	return se
}

func (t *Template) createInputItem(name, label string, required bool, desc ...string) *huh.Input {
	input := huh.NewInput().Key(name).Title(label)
	if required {
		// Validating fields is easy. The form will mark erroneous fields
		// and display error messages accordingly.
		input.Validate(NotBlankValidator(name))
	}
	if len(desc) > 0 && desc[0] != "" {
		input.Description(desc[0])
	}
	return input
}

func (t *Template) createTextAreaItem(name, label string, required bool, desc ...string) *huh.Text {
	// Todo CharLimit??
	text := huh.NewText().Key(name).
		Placeholder("Just put it in the mailbox please").
		Title(label).
		Lines(5)

	if required {
		text.Validate(NotBlankValidator(name))
	}
	if len(desc) > 0 && desc[0] != "" {
		text.Description(desc[0])
	}
	return text
}

// NotBlankValidator is a verification function that checks whether the input is empty
func NotBlankValidator(name string) func(s string) error {
	return func(s string) error {
		if strings.TrimSpace(s) == "" {
			return NewMissingErr(name)
		}
		return nil
	}
}

func ArrayInGroupsOf(fields []huh.Field, num int) []*huh.Group {
	var groups []*huh.Group
	length := len(fields)
	if length <= num {
		groups = append(groups, huh.NewGroup(fields...))
		return groups
	}
	// it should be divided into <quantity> slices
	var quantity int
	if length%num == 0 {
		quantity = length / num
	} else {
		quantity = (length / num) + 1
	}

	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			groups = append(groups, huh.NewGroup(fields[start:end]...))
		} else {
			groups = append(groups, huh.NewGroup(fields[start:]...))
		}
		start = i * num
	}
	return groups
}
