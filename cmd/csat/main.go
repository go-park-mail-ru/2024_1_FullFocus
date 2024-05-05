package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	csatapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/csat/app"
)

const (
	_connTimeout = 5 * time.Second
)

func main() {
	app := csatapp.New()

	go func() {
		app.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	app.GRPCServer.Stop()
}
