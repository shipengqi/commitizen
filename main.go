package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/cmd/cz"
)

const (
	ExitCodeOk        = 0
	ExitCodeException = 1
)

func main() {
	err := cz.New().Execute()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			fmt.Println(err.Error())
			os.Exit(ExitCodeOk)
			return
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(ExitCodeException)
	}
	os.Exit(ExitCodeOk)
}
