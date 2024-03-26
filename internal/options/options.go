package options

import (
	cliflag "github.com/shipengqi/component-base/cli/flag"

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

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GitOptions.AddFlags(fss.FlagSet("Git Commit"))

	fs := fss.FlagSet("Commitizen")
	fs.BoolVar(&o.DryRun, "dry-run", o.DryRun, "do not create a commit, but show the message and list of paths \nthat are to be committed.")
	fs.StringVarP(&o.Template, "template", "t", o.Template, "template name to use when multiple templates exist.")
	fs.BoolVarP(&o.Default, "default", "d", o.Default, "use the default template, '--default' has a higher priority than '--template'.")
	fs.BoolVar(&o.NoTTY, "no-tty", o.NoTTY, "make sure that the TTY (terminal) is never used for any output.")

	_ = fs.MarkHidden("no-tty")

	return
}
