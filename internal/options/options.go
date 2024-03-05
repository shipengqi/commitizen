package options

import (
	"github.com/spf13/pflag"

	"github.com/shipengqi/commitizen/internal/git"
)

type Options struct {
	GitOptions *git.Options
}

func New() *Options {
	return &Options{
		GitOptions: git.NewOptions(),
	}
}

func (o *Options) AddFlags(f *pflag.FlagSet) {
	o.GitOptions.AddFlags(f)
}
