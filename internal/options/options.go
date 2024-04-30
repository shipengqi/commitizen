package options

import (
	"fmt"
	"os"
	"path/filepath"

	cliflag "github.com/shipengqi/component-base/cli/flag"

	"github.com/shipengqi/commitizen/internal/git"
)

type Options struct {
	DryRun     bool
	Default    bool
	Debug      bool
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
	dir := filepath.Join(os.TempDir(), "commitizen/logs")

	o.GitOptions.AddFlags(fss.FlagSet("Git Commit"))

	fs := fss.FlagSet("Commitizen")
	fs.BoolVar(&o.DryRun, "dry-run", o.DryRun, "do not create a commit, but show the message and list of paths \nthat are to be committed.")
	fs.StringVarP(&o.Template, "template", "t", o.Template, "template name to use when multiple templates exist.")
	fs.BoolVarP(&o.Default, "default", "d", o.Default, "use the default template, '--default' has a higher priority than '--template'.")
	fs.BoolVar(&o.Debug, "debug", o.Debug, fmt.Sprintf("enable debug mode, writing log file to the %s directory.", dir))

	_ = fs.MarkHidden("debug")
	return
}
