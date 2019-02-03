package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/irisendozo/stockp-api/internal/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := new(server.Config)
	config.Init()

	// Setup log format
	log.SetFormatter(&log.JSONFormatter{})

	// Setup server
	srv := server.New(server.Options{
		StocksAPIKey: config.StocksAPIKey,
		Port:         config.Port,
	})
	srv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	killSignalReceived := <-c

	log.Fatalf("Process killed with signal: %v", killSignalReceived.String())
	srv.Stop()
}
