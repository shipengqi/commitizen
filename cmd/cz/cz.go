package cz

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shipengqi/commitizen/internal/git"
	"github.com/shipengqi/commitizen/internal/templater"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:  "commitizen",
		Long: `Command line utility to standardize git commit messages.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			isrepo, err := git.IsGitRepo()
			if err != nil {
				return err
			}
			if !isrepo {
				return errors.New("not a git repository")
			}
			tmpl, err := templater.Load([]byte(templater.DefaultCommitTemplate))
			if err != nil {
				return err
			}
			msg, err := tmpl.Run()
			if err != nil {
				return err
			}
			output, err := git.Commit(msg)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	c.SilenceUsage = true
	c.SilenceErrors = true

	c.AddCommand(NewInitCmd())

	return c
}
