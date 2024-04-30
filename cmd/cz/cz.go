package cz

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	cliflag "github.com/shipengqi/component-base/cli/flag"
	"github.com/shipengqi/component-base/term"
	"github.com/shipengqi/golib/convutil"
	"github.com/shipengqi/golib/fsutil"
	"github.com/shipengqi/log"
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
		PreRun: func(cmd *cobra.Command, args []string) {
			if !o.Debug {
				return
			}
			opts := &log.Options{
				DisableRotate:        true,
				DisableFileCaller:    true,
				DisableConsoleCaller: true,
				DisableConsoleLevel:  true,
				DisableConsoleTime:   true,
				Output:               filepath.Join(os.TempDir(), "commitizen/logs"),
				FileLevel:            log.DebugLevel.String(),
				FilenameEncoder:      filenameEncoder,
			}
			err := fsutil.MkDirAll(opts.Output)
			if err != nil {
				panic(err)
			}
			log.Configure(opts)
		},
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

			msg, err := tmpl.Run()
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

	cliflag.InitFlags(c.Flags())

	// applies the FlagSets to this command
	fs := c.Flags()
	fss := o.Flags()
	for _, set := range fss.FlagSets {
		fs.AddFlagSet(set)
	}

	fs.SortFlags = false
	c.DisableFlagsInUseLine = true

	// set both usage and help function.
	width, _, _ := term.TerminalSize(c.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(c, fss, width)

	c.AddCommand(NewInitCmd())
	c.AddCommand(NewVersionCmd())

	return c
}

func filenameEncoder() string {
	return fmt.Sprintf("%s.%s.log", filepath.Base(os.Args[0]), time.Now().Format("20060102150405"))
}
