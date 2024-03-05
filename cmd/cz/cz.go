package cz

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/shipengqi/commitizen/internal/git"
	"github.com/shipengqi/commitizen/internal/render"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:  "commitizen",
		Long: `Command line utility to standardize git commit messages.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			isRepo, err := git.IsGitRepo()
			if err != nil {
				return err
			}
			if !isRepo {
				return errors.New("not a git repository")
			}
			tmpl, err := render.Load([]byte(render.DefaultCommitTemplate))
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
	c.AddCommand(NewLoadCmd())

	return c
}
