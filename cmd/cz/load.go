package cz

import (
	"github.com/spf13/cobra"
)

func NewLoadCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "load",
		Short: "Load templates.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return c
}
