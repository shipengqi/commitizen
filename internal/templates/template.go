package templates

import (
	"bytes"
	standarderrs "errors"
	"fmt"
	"text/template"

	"github.com/charmbracelet/huh"
	"github.com/mitchellh/mapstructure"
	"github.com/shipengqi/golib/strutil"

	"github.com/shipengqi/commitizen/internal/errors"
	"github.com/shipengqi/commitizen/internal/helpers"
	"github.com/shipengqi/commitizen/internal/parameter"
	"github.com/shipengqi/commitizen/internal/parameter/boolean"
	"github.com/shipengqi/commitizen/internal/parameter/integer"
	"github.com/shipengqi/commitizen/internal/parameter/list"
	"github.com/shipengqi/commitizen/internal/parameter/multilist"
	"github.com/shipengqi/commitizen/internal/parameter/secret"
	"github.com/shipengqi/commitizen/internal/parameter/str"
	"github.com/shipengqi/commitizen/internal/parameter/text"
)

type Template struct {
	all    map[parameter.FieldKey]huh.Field
	sorted *SortedGroupMap

	Name    string
	Desc    string
	Format  string
	Default bool
	Items   []map[string]interface{}
	Groups  []*parameter.Group
}

func (t *Template) Initialize() error {
	if strutil.IsEmpty(t.Format) {
		return errors.NewMissingErr("format")
	}

	existGroups := make(map[string]struct{}, len(t.Groups))
	existItems := make(map[string]struct{}, len(t.Items))

	// validate the groups defined in the template
	for _, v := range t.Groups {
		errs := v.Validate()
		if len(errs) > 0 {
			return standarderrs.Join(errs...)
		}

		if _, ok := existGroups[v.Name]; ok {
			return fmt.Errorf("duplicate group name: %s", v.Name)
		}
		existGroups[v.Name] = struct{}{}
	}

	// create a sorted group map
	t.sorted = NewSortedGroupMap(t.Groups...)
	t.all = make(map[parameter.FieldKey]huh.Field)

	for _, v := range t.Items {
		namestr, err := helpers.GetValueFromYAML[string](v, "name")
		if err != nil {
			return err
		}

		if _, ok := existItems[namestr]; ok {
			return fmt.Errorf("duplicate item name: %s", namestr)
		}
		existItems[namestr] = struct{}{}

		typestr, err := helpers.GetValueFromYAML[string](v, "type")
		if err != nil {
			return err
		}
		var (
			param parameter.Interface
			group string
		)
		switch typestr {
		case parameter.TypeBoolean:
			param = &boolean.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeString:
			param = &str.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeInteger:
			param = &integer.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeSecret:
			param = &secret.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeText:
			param = &text.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeList:
			param = &list.Param{}
			err = mapstructure.Decode(v, &param)
		case parameter.TypeMultiList:
			param = &multilist.Param{}
			err = mapstructure.Decode(v, &param)
		default:
			return fmt.Errorf("unknown type %s", typestr)
		}
		if err != nil {
			return err
		}
		errs := param.Validate()
		if len(errs) > 0 {
			return standarderrs.Join(errs...)
		}

		group = param.GetGroup()
		param.Render()

		t.all[parameter.NewFiledKey(namestr, group)] = param

		if fields, ok := t.sorted.GetFields(group); !ok {
			news := make([]huh.Field, 0)
			news = append(news, param)
			t.sorted.SetFields(group, news)
		} else {
			fields = append(fields, param)
			t.sorted.SetFields(group, fields)
		}
	}

	return nil
}

func (t *Template) Run() ([]byte, error) {
	if t.sorted.Length() == 0 {
		return nil, nil
	}

	values := map[string]interface{}{}
	form := huh.NewForm(t.sorted.FormGroups(t.all)...)
	err := form.Run()
	if err != nil {
		return nil, err
	}
	for _, field := range t.all {
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
