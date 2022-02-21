package cz

import (
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialize the 'git cz' and templates.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return c
}
