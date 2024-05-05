package main

import (
	"os"
	"os/signal"
	"syscall"

	profileapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/profile/app"
)

func main() {
	app := profileapp.New()

	go func() {
		app.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	app.GRPCServer.Stop()
}
