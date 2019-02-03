package server

import (
	"github.com/namsral/flag"
)

// Config variables from flags or environment
type Config struct {
	Port         string
	StocksAPIKey string
}

// Init initializes config variables from environment
func (c *Config) Init() {
	flag.StringVar(&c.Port, "port", "30878", "The port for the server to listen on")
	flag.StringVar(&c.StocksAPIKey, "stocks_api_key", "YOMXEEVAWNXO69OT", "The port for the metrics server to listen on")
	flag.Parse()
}
