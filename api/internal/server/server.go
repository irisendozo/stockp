package server

import (
	"context"
	"net/http"
	"time"

	"github.com/irisendozo/stockp-api/internal/app"
	"github.com/irisendozo/stockp-api/internal/pkg/alphavantage"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

// Server defines a HTTP Server
type Server struct {
	*http.Server
}

// Options define the server options passed on to the application and HTTP handlers
type Options struct {
	Port         string
	StocksAPIKey string
}

// New creates a new HTTP server
func New(options Options) Server {
	var (
		// Initialize new application with default router
		stocksAPICaller = alphavantage.New(options.StocksAPIKey)
		router          = NewRouter(app.New(stocksAPICaller))
		addr            = "0.0.0.0:" + options.Port
	)

	return Server{
		&http.Server{
			Addr:         addr,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler: cors.New(cors.Options{
				AllowedOrigins: []string{"http://localhost:3000"},
			}).Handler(router),
		},
	}
}

// Start starts the server listening on a port
func (s Server) Start() {
	log.Printf("Starting server on %s", s.Addr)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("Unable to start server: ", err)
			time.Sleep(time.Second)
		}
	}()
}

// Stop gracefully stops the server
func (s Server) Stop() {
	log.Println("Gracefully stopping server...")
	ctx := context.Background()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server did not stop gracefully.")
	}

	log.Println("Server has stopped.")
}
