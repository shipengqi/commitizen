package v2

import (
	"bytes"
	"errors"
	"fmt"
	errors2 "github.com/shipengqi/commitizen/internal/errors"
	"github.com/shipengqi/golib/strutil"
	"strings"
	"text/template"

	"github.com/charmbracelet/huh"
	"github.com/mitchellh/mapstructure"

	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/boolean"
	"github.com/shipengqi/commitizen/internal/parameter/integer"
	"github.com/shipengqi/commitizen/internal/parameter/list"
	"github.com/shipengqi/commitizen/internal/parameter/multilist"
	"github.com/shipengqi/commitizen/internal/parameter/secret"
	"github.com/shipengqi/commitizen/internal/parameter/str"
	"github.com/shipengqi/commitizen/internal/parameter/text"
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

type Template struct {
	Version string
	Name    string
	Desc    string
	Format  string
	Default bool
	Items   []map[string]interface{}
	groups  []*huh.Group
	fields  []huh.Field
}

func (t *Template) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(t); err != nil {
		return err
	}

	if strutil.IsEmpty(t.Format) {
		return errors2.NewRequiredErr("format")
	}

	groups := make(map[string][]huh.Field)
	exists := make(map[string]struct{}, len(t.Items))
	for _, v := range t.Items {
		namestr, err := GetValueFromYAML[string](v, "name")
		if err != nil {
			return err
		}

		if _, ok := exists[namestr]; ok {
			return fmt.Errorf("duplicate name: %s", namestr)
		}
		exists[namestr] = struct{}{}

		typestr, err := GetValueFromYAML[string](v, "type")
		if err != nil {
			return err
		}
		var (
			param parameter.Interface
			field huh.Field
			group string
		)
		switch typestr {
		case parameter.TypeBoolean:
			param = boolean.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeString:
			param = str.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeInteger:
			param = integer.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeSecret:
			param = secret.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeText:
			param = text.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeList:
			param = list.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeMultiList:
			param = multilist.Param{}
			err = mapstructure.Decode(v, &param)
		default:
			return fmt.Errorf("unknown type %s", typestr)
		}
		if err != nil {
			return err
		}
		group = param.GetGroup()
		field = param.Render()
		t.fields = append(t.fields, field)
		if group == "" {
			group = "unknown"
		}
		if fields, ok := groups[group]; !ok {
			news := make([]huh.Field, 0)
			news = append(news, field)
			groups[group] = news
		} else {
			fields = append(fields, field)
			groups[group] = fields
		}
	}

	t.groups = GroupMap2Array(groups)
	return nil
}

func (t *Template) Run() ([]byte, error) {

	if len(t.groups) == 0 {
		return nil, nil
	}

	values := map[string]interface{}{}
	form := huh.NewForm(t.groups...)
	err := form.Run()
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

// ---------------------------------
// Helpers

func GetValueFromYAML[T any](data map[string]interface{}, key string) (T, error) {
	var (
		res T
		ok  bool
		v   interface{}
	)

	v, ok = data[key]
	if !ok {
		return res, errors2.NewRequiredErr(key)
	}
	res, ok = v.(T)
	if !ok {
		return res, errors.New("error type")
	}
	return res, nil
}

func GroupMap2Array(all map[string][]huh.Field) []*huh.Group {
	var groups []*huh.Group
	for _, v := range all {
		groups = append(groups, huh.NewGroup(v...))
	}

	return groups
}
