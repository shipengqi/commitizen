package git

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Quiet   bool
	Verbose bool
	SignOff bool
	All     bool
	Amend   bool
	DryRun  bool
	Author  string
	Date    string
}

func NewOptions() *Options {
	return &Options{
		Quiet:   false,
		Verbose: false,
		SignOff: false,
		All:     false,
		Amend:   false,
		DryRun:  false,
		Author:  "",
		Date:    "",
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	// inherits the --dry-run argument from the parent command
	fs.BoolVarP(&o.Quiet, "quiet", "q", o.Quiet, "suppress summary after successful commit")
	fs.BoolVarP(&o.Verbose, "verbose", "v", o.Verbose, "show diff in commit message template")
	fs.StringVar(&o.Author, "author", o.Author, "override author for commit")
	fs.StringVar(&o.Date, "date", o.Date, "override date for commit")
	fs.BoolVarP(&o.All, "all", "a", o.All, "commit all changed files.")
	fs.BoolVarP(&o.SignOff, "signoff", "s", o.SignOff, "add a Signed-off-by trailer.")
	fs.BoolVar(&o.Amend, "amend", o.Amend, "amend previous commit")
	return
}

func (o *Options) Combine(filename string) []string {
	combination := []string{
		"commit",
		"-F",
		filename,
	}
	if o.Quiet {
		combination = append(combination, "-q")
	}
	if o.Verbose {
		combination = append(combination, "-v")
	}
	if o.Author != "" {
		combination = append(combination, "--author", o.Author)
	}
	if o.Date != "" {
		combination = append(combination, "--date", o.Date)
	}
	if o.All {
		combination = append(combination, "-a")
	}
	if o.Amend {
		combination = append(combination, "--amend")
	}
	if o.DryRun {
		combination = append(combination, "--dry-run")
	}

	return combination
}
