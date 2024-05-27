package main

import (
	"os"
	"os/signal"
	"syscall"

	promotionapp "github.com/go-park-mail-ru/2024_1_FullFocus/microservices/promotion/app"
)

func main() {
	app := promotionapp.New()

	go func() {
		app.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	app.GRPCServer.Stop()
}
