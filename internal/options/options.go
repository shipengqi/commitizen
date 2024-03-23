package options

import (
	"github.com/spf13/pflag"

	"github.com/shipengqi/commitizen/internal/git"
)

type Options struct {
	DryRun     bool
	NoTTY      bool
	Default    bool
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

	f.BoolVar(&o.DryRun, "dry-run", o.DryRun, "do not create a commit, but show the message and list of paths \nthat are to be committed.")
	f.StringVarP(&o.Template, "template", "t", o.Template, "template name to use when multiple templates exist.")
	f.BoolVarP(&o.Default, "default", "d", o.Default, "use the default template, '--default' has a higher priority than '--template'.")
	f.BoolVar(&o.NoTTY, "no-tty", o.NoTTY, "make sure that the TTY (terminal) is never used for any output.")

	_ = f.MarkHidden("no-tty")
}
