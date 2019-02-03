package server_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/irisendozo/stockp-api/internal/app"
	"github.com/irisendozo/stockp-api/internal/pkg/alphavantage"
	"github.com/irisendozo/stockp-api/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestFetchBalanceEndpoint(t *testing.T) {
	s, _ := serverRouterSetup(t)

	req, _ := http.NewRequest("GET", "/balance/me", nil)
	response, code := executeRequest(s, req)

	var balanceResp app.BalanceAPIResponse
	if err := json.Unmarshal(response, &balanceResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, app.BalanceAPIResponse{
		Balance: 0,
	}, balanceResp)
	assert.Equal(t, 200, code)
}

func TestGetHistoryEndpoint(t *testing.T) {
	s, mock := serverRouterSetup(t)

	req, _ := http.NewRequest("POST", "/stocks/buy/any", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("GET", "/stocks/history/me", nil)
	response, code := executeRequest(s, req)

	var historyResp []app.Stock
	if err := json.Unmarshal(response, &historyResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, []app.Stock{
		{
			Symbol: "somesymbol",
			Name:   "somename",
			Count:  1,
			Price:  1,
		},
	}, historyResp)
	assert.Equal(t, 200, code)
}

func TestAddBalanceEndpoint(t *testing.T) {
	s, _ := serverRouterSetup(t)

	req, _ := http.NewRequest("POST", "/balance/add/1", nil)
	response, code := executeRequest(s, req)

	var balanceResp app.BalanceAPIResponse
	if err := json.Unmarshal(response, &balanceResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, app.BalanceAPIResponse{
		Balance: 1,
	}, balanceResp)
	assert.Equal(t, 200, code)
}

func TestWithdrawBalanceEndpoint(t *testing.T) {
	s, _ := serverRouterSetup(t)

	req, _ := http.NewRequest("POST", "/balance/add/5", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/balance/withdraw/1", nil)
	response, code := executeRequest(s, req)

	var balanceResp app.BalanceAPIResponse
	if err := json.Unmarshal(response, &balanceResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, app.BalanceAPIResponse{
		Balance: 4,
	}, balanceResp)
	assert.Equal(t, 200, code)
}

func TestWithdrawBalanceEndpoint_InsufficientFunds(t *testing.T) {
	s, _ := serverRouterSetup(t)

	req, _ := http.NewRequest("POST", "/balance/withdraw/1", nil)
	response, code := executeRequest(s, req)

	var balanceResp app.BalanceAPIResponse
	if err := json.Unmarshal(response, &balanceResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.BalanceAPIResponse{
		Balance: 0,
	}, balanceResp)
	assert.Equal(t, 500, code)
}

func TestFetchOwnedStocksEndpoint(t *testing.T) {
	s, _ := serverRouterSetup(t)

	req, _ := http.NewRequest("GET", "/stocks/me", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, app.StocksAPIResponse{
		OwnedStocks: []app.Stock{},
	}, stocksResp)
	assert.Equal(t, 200, code)
}

func TestFetchSearchedStocksEndpoint(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{
		Symbol: "NKE",
		Name:   "NIKE Inc.",
		Price:  80.33,
	}, nil).AnyTimes()

	req, _ := http.NewRequest("GET", "/stocks/search/NKE", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.SearchStockAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, app.SearchStockAPIResponse{
		Symbol: "NKE",
		Name:   "NIKE Inc.",
		Price:  80.33,
	}, stocksResp)
	assert.Equal(t, 200, code)
}

func TestFetchSearchedStocksEndpoint_FailureFromAlphavantage(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().
		GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{}, errors.New("api limit")).AnyTimes()

	req, _ := http.NewRequest("GET", "/stocks/search/NKE", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.SearchStockAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.SearchStockAPIResponse{}, stocksResp)
	assert.Equal(t, 500, code)
}

func TestBuyStockEndpoint(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{
		Symbol:       "NKE",
		Name:         "NIKE Inc.",
		Price:        80.33,
		PurchaseDate: mock.Anything,
	}, nil).AnyTimes()

	req, _ := http.NewRequest("POST", "/balance/add/10000", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/buy/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, 1, stocksResp.OwnedStocks[0].Count)
	assert.Equal(t, "NKE", stocksResp.OwnedStocks[0].Symbol)
	assert.Equal(t, 200, code)
}

func TestBuyStockEndpoint_InsufficientFunds(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{
		Symbol: "NKE",
		Name:   "NIKE Inc.",
		Price:  80.33,
	}, nil).AnyTimes()

	req, _ := http.NewRequest("POST", "/stocks/buy/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.StocksAPIResponse{}, stocksResp)
	assert.Equal(t, 500, code)
}

func TestBuyStockEndpoint_FailureFromAlphavantage(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().
		GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{}, errors.New("api limit")).AnyTimes()

	req, _ := http.NewRequest("POST", "/balance/add/10000", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/buy/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.StocksAPIResponse{}, stocksResp)
	assert.Equal(t, 500, code)
}

func TestSellStockEndpoint(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{
		Symbol: "NKE",
		Name:   "NIKE Inc.",
		Price:  80.33,
	}, nil).AnyTimes()

	req, _ := http.NewRequest("POST", "/balance/add/10000", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/buy/NKE/5", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/sell/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		t.Fatalf("failed to unmarshal: %s", err)
	}

	assert.Equal(t, 4, stocksResp.OwnedStocks[0].Count)
	assert.Equal(t, "NKE", stocksResp.OwnedStocks[0].Symbol)
	assert.Equal(t, 200, code)
}

func TestSellStockEndpoint_InsufficientStocks(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{
		Symbol: "NKE",
		Name:   "NIKE Inc.",
		Price:  80.33,
	}, nil).AnyTimes()

	req, _ := http.NewRequest("POST", "/stocks/sell/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.StocksAPIResponse{}, stocksResp)
	assert.Equal(t, 500, code)
}

func TestSellStockEndpoint_FailureFromAlphavantage(t *testing.T) {
	s, stockAPICaller := serverRouterSetup(t)

	stockAPICaller.EXPECT().
		GetLatestPriceForStock("NKE").Return(&alphavantage.LatestStockPrice{}, errors.New("api limit")).AnyTimes()

	req, _ := http.NewRequest("POST", "/balance/add/10000", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/buy/NKE/5", nil)
	executeRequest(s, req)

	req, _ = http.NewRequest("POST", "/stocks/sell/NKE/1", nil)
	response, code := executeRequest(s, req)

	var stocksResp app.StocksAPIResponse
	if err := json.Unmarshal(response, &stocksResp); err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, app.StocksAPIResponse{}, stocksResp)
	assert.Equal(t, 500, code)
}

func serverRouterSetup(t *testing.T) (*mux.Router, *alphavantage.MockCaller) {
	var (
		ctrl   = gomock.NewController(t)
		m      = alphavantage.NewMockCaller(ctrl)
		router = server.NewRouter(app.New(m))
	)

	return router, m
}

func executeRequest(router *mux.Router, req *http.Request) ([]byte, int) {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	responseData, _ := ioutil.ReadAll(rr.Body)

	return responseData, rr.Code
}
