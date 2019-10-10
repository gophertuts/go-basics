package main

import (
	"log"
	"os"
	"syscall"

	"github.com/gophertuts/go-basics/package-management/basic-http/app"
)

func main() {
	application := app.New()
	go func() {
		err := application.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	app.ListenForSignals([]os.Signal{syscall.SIGINT, syscall.SIGTERM}, application)
}
