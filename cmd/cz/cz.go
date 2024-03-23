package cz

import (
	"errors"
	"fmt"

	"github.com/shipengqi/golib/convutil"
	"github.com/spf13/cobra"

	"github.com/shipengqi/commitizen/internal/config"
	"github.com/shipengqi/commitizen/internal/git"
	"github.com/shipengqi/commitizen/internal/options"
)

func New() *cobra.Command {
	o := options.New()
	c := &cobra.Command{
		Use:  "commitizen",
		Long: `Command line utility to standardize git commit messages.`,
		RunE: func(_ *cobra.Command, _ []string) error {
			isRepo, err := git.IsGitRepo()
			if err != nil {
				return err
			}
			if !isRepo {
				return errors.New("not a git repository")
			}

			conf := config.New()
			tmpl, err := conf.Run(o)
			if err != nil {
				return err
			}

			msg, err := tmpl.Run(o.NoTTY)
			if err != nil {
				return err
			}

			if o.DryRun {
				fmt.Println(convutil.B2S(msg))
				fmt.Println("")
				// inherits the --dry-run argument from the parent command
				o.GitOptions.DryRun = o.DryRun
			}
			output, err := git.Commit(msg, o.GitOptions)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	c.SilenceUsage = true
	c.SilenceErrors = true
	cobra.EnableCommandSorting = false
	c.CompletionOptions.DisableDefaultCmd = true

	f := c.Flags()
	f.SortFlags = false
	c.DisableFlagsInUseLine = true
	o.AddFlags(f)

	c.AddCommand(NewInitCmd())
	c.AddCommand(NewVersionCmd())

	return c
}
