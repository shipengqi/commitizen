package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/shipengqi/commitizen/cmd/cz"
	"github.com/shipengqi/commitizen/internal/ui"
	"os"
)

const (
	ExitCodeOk        = 0
	ExitCodeException = 1
	ExitCodeSignal    = 2
)

func main() {
	// err := execute()
	// if err != nil {
	// 	if err == terminal.InterruptErr {
	// 		os.Exit(ExitCodeSignal)
	// 	}
	// 	fmt.Printf("exception: %s\n", err)
	// 	os.Exit(ExitCodeException)
	// }
	// os.Exit(ExitCodeOk)

	i := ui.NewInput("Scope. Could be anything specifying place of the commit change:")
	if _, err := tea.NewProgram(i).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}

func execute() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[recover panic]:\n%s\n", err)
		}
	}()

	return cz.New().Execute()
}
