package render

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	TypeSelect    = "select"
	TypeInput     = "input"
	TypeTextArea  = "textarea"
	DescMaxLength = 50
)

func LoadTemplates(files ...string) ([]*Template, error) {
	if len(files) == 0 {
		return nil, nil
	}

	var templates []*Template
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		tmpl, err := Load(data)
		if err != nil {
			return nil, err
		}
		templates = append(templates, tmpl)
	}
	return templates, nil
}

func Load(data []byte) (*Template, error) {
	var tmpl Template

	err := yaml.Unmarshal(data, &tmpl)
	if err != nil {
		return nil, err
	}

	return &tmpl, nil
}

// -----------------------------------------

func isEmptyStr(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}
