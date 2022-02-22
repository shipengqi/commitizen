package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/shipengqi/commitizen/cmd/cz"
)

const (
	ExitCodeOk        = 0
	ExitCodeException = 1
	ExitCodeSignal    = 2
)

func main() {
	err := execute()
	if err != nil {
		if err == terminal.InterruptErr {
			os.Exit(ExitCodeSignal)
		}
		fmt.Printf("exception: %s\n", err)
		os.Exit(ExitCodeException)
	}
	os.Exit(ExitCodeOk)
}

func execute() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[recover panic]:\n%s\n", err)
		}
	}()

	return cz.New().Execute()
}
