package cz

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/shipengqi/commitizen/internal/git"
)

func NewInitCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "init",
		Short: "Initialize the 'git cz' and templates.",
		RunE: func(cmd *cobra.Command, args []string) error {
			src, err := exec.LookPath(os.Args[0])
			if err != nil {
				return err
			}
			dst, err := git.InitSubCmd(src, "cz")
			if err != nil {
				return err
			}
			fmt.Printf("Init commitizen to %s\n", dst)
			return nil
		},
	}

	return c
}
