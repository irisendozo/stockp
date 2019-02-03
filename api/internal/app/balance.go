package app

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// BalanceAPIResponse defines the API response object for balance API
type BalanceAPIResponse struct {
	Balance float64
}

// FetchBalance handles request for getting current balance
func (a *Application) FetchBalance(w http.ResponseWriter, r *http.Request) {
	a.GenerateBalanceAPIResponse(w)
}

// AddBalance handles request for adding balance
func (a *Application) AddBalance(w http.ResponseWriter, r *http.Request) {
	var (
		amount = a.ConvertStringToFloat(mux.Vars(r)["amount"])
	)

	a.BalanceHistory = append(a.BalanceHistory, amount)

	a.GenerateBalanceAPIResponse(w)
}

// WithdrawBalance handles request for withdrawing balance
func (a *Application) WithdrawBalance(w http.ResponseWriter, r *http.Request) {
	var (
		amount = a.ConvertStringToFloat(mux.Vars(r)["amount"])
	)

	if amount > a.GetTotalBalance() {
		err := errors.New("insufficient funds")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.BalanceHistory = append(a.BalanceHistory, -amount)

	a.GenerateBalanceAPIResponse(w)
}

// generateBalanceAPIResponse generates the generic JSON response for the API containing the current balance
func (a *Application) GenerateBalanceAPIResponse(w http.ResponseWriter) {
	response := BalanceAPIResponse{
		Balance: a.GetTotalBalance(),
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

// GetTotalBalance computes the balance from the balance history
func (a *Application) GetTotalBalance() float64 {
	var totalBalance float64
	for _, balance := range a.BalanceHistory {
		totalBalance += balance
	}

	return totalBalance
}
