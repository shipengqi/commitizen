package cz

import (
	"fmt"

	"github.com/shipengqi/component-base/version"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print the version information.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Get().String())
		},
	}
	return c
}
