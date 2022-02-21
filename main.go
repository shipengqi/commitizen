package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shipengqi/commitizen/cmd/cz"
)

const (
	ExitCodeOk        = 0
	ExitCodeException = 1
	ExitCodeSignal    = 2
)

func main() {


	go watch()
	err := execute()
	if err != nil {
		log.Printf("exception: %s", err)
		os.Exit(ExitCodeException)
	}
	os.Exit(ExitCodeOk)
}

func execute() error {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[recover panic]:\n%s", err)
		}
	}()

	return cz.New().Execute()
}

func watch() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case sig := <-quit:
			log.Printf("get a signal %s", sig.String())
			os.Exit(ExitCodeSignal)
			return
		}
	}
}
