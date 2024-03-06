package options

import (
	"github.com/spf13/pflag"

	"github.com/shipengqi/commitizen/internal/git"
)

type Options struct {
	DryRun     bool
	GitOptions *git.Options
}

func New() *Options {
	return &Options{
		GitOptions: git.NewOptions(),
		DryRun:     false,
	}
}

func (o *Options) AddFlags(f *pflag.FlagSet) {
	o.GitOptions.AddFlags(f)

	f.BoolVar(&o.DryRun, "dry-run", o.DryRun, "you can use the --dry-run flag to preview the message that would be committed, without really submitting it.")
}
