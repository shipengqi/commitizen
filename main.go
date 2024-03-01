package main

import (
	"errors"
	"fmt"
	"github.com/shipengqi/commitizen/internal/render"
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
		if errors.Is(err, render.ErrCanceled) {
			fmt.Println(err.Error())
			os.Exit(ExitCodeOk)
			return
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(ExitCodeException)
	}
	os.Exit(ExitCodeOk)
}
