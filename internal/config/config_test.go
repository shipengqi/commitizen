package config_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/shipengqi/commitizen/internal/config"
)

func TestLoadTemplates(t *testing.T) {
	tests := []struct {
		title  string
		tmpl   string
		expect int
	}{
		{
			"without any default templates",
			"./testdata/.git-czrc",
			2,
		},
		{
			"with one default templates",
			"./testdata/.git-czrc-with-default",
			2,
		},
		{
			"with two default templates",
			"./testdata/.git-czrc-with-two-default",
			2,
		},
	}
	g := NewWithT(t)
	for _, v := range tests {
		t.Run(v.title, func(t *testing.T) {
			ts, err := config.LoadTemplates(v.tmpl)
			g.Expect(err).To(BeNil())
			g.Expect(len(ts)).To(Equal(v.expect))
		})
	}
}
