package options

import (
	"github.com/spf13/pflag"

	"github.com/shipengqi/commitizen/internal/git"
)

type Options struct {
	DryRun     bool
	NoTTY      bool
	Template   string
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
	f.StringVarP(&o.Template, "template", "t", o.Template, "template name to use when multiple templates exist.")
	f.BoolVar(&o.NoTTY, "no-tty", o.NoTTY, "make sure that the TTY (terminal) is never used for any output.")

	_ = f.MarkHidden("no-tty")
}
