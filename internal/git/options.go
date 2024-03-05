package git

import "github.com/spf13/pflag"

type Options struct {
	SignOff bool
	Add     bool
}

func NewOptions() *Options {
	return &Options{
		SignOff: false,
		Add:     false,
	}
}

func (o *Options) AddFlags(f *pflag.FlagSet) {
	f.BoolVarP(&o.SignOff, "signoff", "s", o.SignOff, "Add a Signed-off-by trailer by the committer at the end of the commit log message.")
	f.BoolVarP(&o.Add, "add", "a", o.Add, "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.")
}
