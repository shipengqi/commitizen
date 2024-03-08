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
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(version.Get().String())
		},
	}
	return c
}
