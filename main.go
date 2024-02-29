package main

import (
	"fmt"
	"os"

	"github.com/shipengqi/commitizen/cmd/cz"
)

const (
	ExitCodeOk        = 0
	ExitCodeException = 1
)

func main() {
	err := cz.New().Execute()
	if err != nil {
		fmt.Printf("exception: %s\n", err)
		os.Exit(ExitCodeException)
	}
	os.Exit(ExitCodeOk)
}
