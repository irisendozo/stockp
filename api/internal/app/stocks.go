package app

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// StocksAPIResponse defines the API response object for stocks API
type StocksAPIResponse struct {
	OwnedStocks []Stock
}

// SearchStockAPIResponse defines the API response object for search stocks API
type SearchStockAPIResponse struct {
	Symbol string
	Name   string
	Price  float64
}

func (a *Application) FetchPurchaseStockHistory(w http.ResponseWriter, r *http.Request) {
	response := a.PurchaseStockHistory

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// FetchOwnedStock handles request for getting currently owned stocks
func (a *Application) FetchOwnedStocks(w http.ResponseWriter, r *http.Request) {
	a.GenerateStocksAPIResponse(w)
}

// FetchSearchedStocks handles request for searching stocks
func (a *Application) FetchSearchedStocks(w http.ResponseWriter, r *http.Request) {
	var (
		filterParam = mux.Vars(r)["filter"]
	)

	matchingStock, err := a.StocksAPICaller.GetLatestPriceForStock(filterParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := SearchStockAPIResponse{
		Symbol: matchingStock.Symbol,
		Name:   matchingStock.Name,
		Price:  matchingStock.Price,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// BuyStock handles request for checking balance if user can buy stock price * quantity and
// if successful, updates balance and adds stock to currently owned stock
func (a *Application) BuyStock(w http.ResponseWriter, r *http.Request) {
	var (
		symbolParam   = mux.Vars(r)["symbol"]
		quantityParam = a.ConvertStringToFloat(mux.Vars(r)["quantity"])
	)

	stock, err := a.StocksAPICaller.GetLatestPriceForStock(symbolParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	priceOfRequest := quantityParam * stock.Price
	if priceOfRequest < a.GetTotalBalance() {
		a.PurchaseStockHistory = append(a.PurchaseStockHistory, Stock{
			Symbol:       stock.Symbol,
			Name:         stock.Name,
			Count:        int(quantityParam),
			Price:        stock.Price,
			PurchaseDate: time.Now(),
		})
		a.BalanceHistory = append(a.BalanceHistory, -priceOfRequest)
	} else {
		http.Error(w, "Insufficient funds to buy stock", http.StatusInternalServerError)
		return
	}

	a.GenerateStocksAPIResponse(w)
}

// SellStock handles request for checking if user has enough stock quantity and
// if successful, updates balance and minuses stock to currently owned stock
func (a *Application) SellStock(w http.ResponseWriter, r *http.Request) {
	var (
		symbolParam   = mux.Vars(r)["symbol"]
		quantityParam = a.ConvertStringToFloat(mux.Vars(r)["quantity"])
	)

	stock, err := a.StocksAPICaller.GetLatestPriceForStock(symbolParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	priceOfRequest := quantityParam * stock.Price
	if ownedStock := a.GetOwnedStock(stock.Symbol); ownedStock != nil && ownedStock.Count > int(quantityParam) {
		a.PurchaseStockHistory = append(a.PurchaseStockHistory, Stock{
			Symbol: stock.Symbol,
			Name:   stock.Name,
			Count:  -int(quantityParam),
			Price:  stock.Price,
		})
		a.BalanceHistory = append(a.BalanceHistory, priceOfRequest)
	} else {
		http.Error(w, "Insufficient stocks to sell", http.StatusInternalServerError)
		return
	}

	a.GenerateStocksAPIResponse(w)
}

// GenerateStocksAPIResponse generates the generic JSON response for the API containing the current balance
func (a *Application) GenerateStocksAPIResponse(w http.ResponseWriter) {
	response := StocksAPIResponse{
		OwnedStocks: a.GetAllOwnedStocks(),
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllOwnedStocks meshes purchase history to create list of stock purchases
func (a *Application) GetAllOwnedStocks() []Stock {
	ownedStocks := []Stock{}

	symbols := []string{}
	for _, purchasedStock := range a.PurchaseStockHistory {
		symbols = append(symbols, purchasedStock.Symbol)
	}

	cleanedSymbolList := a.RemoveDuplicateSymbols(symbols)

	for _, symbol := range cleanedSymbolList {
		if ownedStock := a.GetOwnedStock(symbol); ownedStock != nil {
			ownedStocks = append(ownedStocks, *ownedStock)
		}
	}

	return ownedStocks
}

// GetOwnedStock meshes purchase history to create list of stock purchases for one stock
func (a *Application) GetOwnedStock(symbol string) *Stock {
	var count int
	var name string
	var price float64
	var purchaseDate time.Time

	for _, purchasedStock := range a.PurchaseStockHistory {
		if purchasedStock.Symbol == symbol {
			count += purchasedStock.Count
			price += purchasedStock.Price
			name = purchasedStock.Name
			purchaseDate = purchasedStock.PurchaseDate
		}
	}

	if count == 0 {
		return nil
	}

	return &Stock{
		Symbol:       symbol,
		Name:         name,
		Count:        count,
		PurchaseDate: purchaseDate,
		Price:        math.Round(price / float64(count)),
	}
}

func (a *Application) RemoveDuplicateSymbols(symbols []string) []string {
	result := []string{}

	for i := 0; i < len(symbols); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if symbols[v] == symbols[i] {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, symbols[i])
		}
	}
	return result
}
