package templater

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const (
	TypeSelect    = "select"
	TypeInput     = "input"
	TypeMultiline = "multiline"
	DescMaxLength = 50
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
	Options  []*Option
	Required bool
}

type Template struct {
	Name   string
	Format string
	Items  []*Item
	// questions []*survey.Question
}

func (t *Template) Run() ([]byte, error) {
	if err := t.init(); err != nil {
		return nil, err
	}

	fmt.Println("All commit message lines will be cropped at 100 characters.")
	// ask the question
	// answers := map[string]interface{}{}
	// if err := survey.Ask(t.questions, &answers); err != nil {
	// 	return nil, err
	// }
	//
	// tmpl, err := template.New("commmitizen").Parse(t.Format)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// for k, v := range answers {
	// 	if option, ok := v.(survey.OptionAnswer); ok {
	// 		answers[k] = option.Value
	// 	} else if vstr, ok := v.(string); ok {
	// 		answers[k] = strings.TrimSpace(vstr)
	// 	}
	// }
	//
	// var buf bytes.Buffer
	// if err = tmpl.Execute(&buf, answers); err != nil {
	// 	return nil, err
	// }
	//
	// return buf.Bytes(), nil
	return nil, nil
}

func (t *Template) init() error {
	if len(t.Format) == 0 {
		return errors.New("format is required")
	}

	for _, item := range t.Items {
		if len(item.Name) == 0 {
			return errors.New("item.name is required")
		}
		if len(item.Desc) == 0 {
			return errors.New("item.desc is required")
		}
		if len(item.Type) == 0 {
			return errors.New("item.type is required")
		}
		// q := survey.Question{
		// 	Name: item.Name,
		// }

		// switch item.Type {
		// case TypeInput:
		// 	q.Prompt = &survey.Input{
		// 		Message: item.Desc,
		// 	}
		// 	if item.Required {
		// 		q.Validate = survey.ComposeValidators(survey.Required, survey.MaxLength(DescMaxLength))
		// 	}
		// case TypeSelect:
		// 	if len(item.Options) == 0 {
		// 		return errors.New("item.options is required in the 'select'")
		// 	}
		// 	s := &survey.Select{
		// 		Message: item.Desc,
		// 	}
		// 	for _, option := range item.Options {
		// 		s.Options = append(s.Options, option.String())
		// 	}
		// 	q.Prompt = s
		// 	q.Transform = func(options []*Option) func(interface{}) interface{} {
		// 		return func(ans interface{}) (newAns interface{}) {
		// 			if answer, ok := ans.(survey.OptionAnswer); !ok {
		// 				return nil
		// 			} else {
		// 				answer.Value = options[answer.Index].Name
		// 				return answer
		// 			}
		// 		}
		// 	}(item.Options)
		// 	if item.Required {
		// 		q.Validate = survey.ComposeValidators(survey.Required, survey.MaxLength(DescMaxLength))
		// 	}
		// case TypeMultiline:
		// 	q.Prompt = &survey.Multiline{
		// 		Message: item.Desc,
		// 	}
		// 	if item.Required {
		// 		q.Validate = survey.Required
		// 	}
		// default:
		// 	continue
		// }
		//
		// t.questions = append(t.questions, &q)
	}
	return nil
}

func LoadFiles(files ...string) ([]*Template, error) {
	if len(files) == 0 {
		return nil, nil
	}

	var tmpls []*Template
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		tmpl, err := Load(data)
		if err != nil {
			return nil, err
		}
		tmpls = append(tmpls, tmpl)
	}
	return tmpls, nil
}

func Load(data []byte) (*Template, error) {
	var tmpl Template

	err := yaml.Unmarshal(data, &tmpl)
	if err != nil {
		return nil, err
	}

	return &tmpl, nil
}
