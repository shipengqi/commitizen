package cz

import (
	"log"

	"github.com/shipengqi/commitizen/internal/templater"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	c := &cobra.Command{
		Use:  "commitizen",
		Long: `Command line utility to standardize git commit messages.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tmpl, err := templater.Load([]byte(templater.DefaultCommitTemplate))
			if err != nil {
				return err
			}
			msg, err := tmpl.Run()
			if err != nil {
				return err
			}
			log.Print(string(msg))
			return nil
		},
	}

	c.SilenceUsage = true
	c.SilenceErrors = true

	return c
}
