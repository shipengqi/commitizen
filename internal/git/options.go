package git

import (
	"strings"

	"github.com/spf13/pflag"
)

type Options struct {
	Quiet         bool
	Verbose       bool
	SignOff       bool
	All           bool
	Amend         bool
	DryRun        bool
	NoVerify      bool
	Author        string
	Date          string
	ExtraGitFlags []string
}

func NewOptions() *Options {
	return &Options{
		Quiet:         false,
		Verbose:       false,
		SignOff:       false,
		All:           false,
		Amend:         false,
		NoVerify:      false,
		DryRun:        false,
		Author:        "",
		Date:          "",
		ExtraGitFlags: []string{},
	}
}

func (o *Options) AddFlags(f *pflag.FlagSet) {
	// inherits the --dry-run argument from the parent command
	f.BoolVarP(&o.Quiet, "quiet", "q", o.Quiet, "suppress summary after successful commit")
	f.BoolVarP(&o.Verbose, "verbose", "v", o.Verbose, "show diff in commit message template")
	f.StringVar(&o.Author, "author", o.Author, "override author for commit")
	f.StringVar(&o.Date, "date", o.Date, "override date for commit")
	f.BoolVarP(&o.All, "all", "a", o.All, "commit all changed files.")
	f.BoolVarP(&o.SignOff, "signoff", "s", o.SignOff, "add a Signed-off-by trailer.")
	f.BoolVar(&o.Amend, "amend", o.Amend, "amend previous commit")
	f.BoolVarP(&o.NoVerify, "no-verify", "n", o.NoVerify, "bypass pre-commit and commit-msg hooks.")
	f.StringSliceVar(&o.ExtraGitFlags, "git-flag", o.ExtraGitFlags, "git flags, e.g. --git-flag=\"--branch\"")
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
	if o.NoVerify {
		combination = append(combination, "--no-verify")
	}
	if o.DryRun {
		combination = append(combination, "--dry-run")
	}
	if len(o.ExtraGitFlags) > 0 {
		result := deDuplicateFlag(o.ExtraGitFlags, "-F", "--file")
		combination = append(combination, result...)
	}

	return combination
}

func deDuplicateFlag(sli []string, short, long string) []string {
	var result []string
	for _, s := range sli {
		if strings.HasPrefix(s, short) {
			continue
		}
		if strings.HasPrefix(s, long) {
			continue
		}
		result = append(result, s)
	}
	return result
}
