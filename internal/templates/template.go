package templates

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/charmbracelet/huh"
	"github.com/mitchellh/mapstructure"
	"github.com/shipengqi/golib/strutil"

	"github.com/shipengqi/commitizen/internal/errors"
	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/boolean"
	"github.com/shipengqi/commitizen/internal/parameter/integer"
	"github.com/shipengqi/commitizen/internal/parameter/list"
	"github.com/shipengqi/commitizen/internal/parameter/multilist"
	"github.com/shipengqi/commitizen/internal/parameter/secret"
	"github.com/shipengqi/commitizen/internal/parameter/str"
	"github.com/shipengqi/commitizen/internal/parameter/text"
)

const UnknownGroup = "unknown"

type Template struct {
	Name    string
	Desc    string
	Format  string
	Default bool
	Items   []map[string]interface{}
	groups  []*huh.Group
	fields  []huh.Field
}

func (t *Template) Initialize() error {
	if strutil.IsEmpty(t.Format) {
		return errors.NewRequiredErr("format")
	}

	groups := NewSortedGroupMap()
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
			group = UnknownGroup
		}
		if fields, ok := groups.Get(group); !ok {
			news := make([]huh.Field, 0)
			news = append(news, field)
			groups.Set(group, news)
		} else {
			fields = append(fields, field)
			groups.Set(group, fields)
		}
	}

	t.groups = groups.FormGroups()
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
		return res, errors.NewRequiredErr(key)
	}
	res, ok = v.(T)
	if !ok {
		return res, errors.ErrType
	}
	return res, nil
}
