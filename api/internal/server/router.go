package server

import (
	"github.com/gorilla/mux"
	"github.com/irisendozo/stockp-api/internal/app"
)

// NewRouter defines a new gorilla MUX router with the server-defined routes
// Add new API routes here
func NewRouter(a *app.Application) *mux.Router {
	var router = mux.NewRouter()

	// Stocks-related routes
	router.HandleFunc("/stocks/me", a.FetchOwnedStocks).Methods("GET")
	router.HandleFunc("/stocks/search/{filter}", a.FetchSearchedStocks).Methods("GET")
	router.HandleFunc("/stocks/buy/{symbol}/{quantity}", a.BuyStock).Methods("POST")
	router.HandleFunc("/stocks/sell/{symbol}/{quantity}", a.SellStock).Methods("POST")
	router.HandleFunc("/stocks/history/me", a.FetchPurchaseStockHistory).Methods("GET")

	// Balance-related routes
	router.HandleFunc("/balance/me", a.FetchBalance).Methods("GET")
	router.HandleFunc("/balance/add/{amount}", a.AddBalance).Methods("POST")
	router.HandleFunc("/balance/withdraw/{amount}", a.WithdrawBalance).Methods("POST")

	return router
}
