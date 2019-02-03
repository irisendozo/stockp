package app

import (
	"strconv"
	"time"

	"github.com/irisendozo/stockp-api/internal/pkg/alphavantage"
)

// Stock is unit local implementation of the currently owned stocks that have been bought by the user
type Stock struct {
	Symbol       string
	Name         string
	Count        int
	Price        float64
	PurchaseDate time.Time
}

// Application implements all REST interface methods for the server
type Application struct {
	StocksAPICaller      alphavantage.Caller
	PurchaseStockHistory []Stock
	BalanceHistory       []float64
}

// New creates a new application initialized with mux router
func New(stocksAPICaller alphavantage.Caller) *Application {
	return &Application{
		StocksAPICaller:      stocksAPICaller,
		PurchaseStockHistory: []Stock{},
		BalanceHistory:       []float64{},
	}
}

// ConvertStringToFloat converts string to float with a default to 0
func (a *Application) ConvertStringToFloat(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return num
}
