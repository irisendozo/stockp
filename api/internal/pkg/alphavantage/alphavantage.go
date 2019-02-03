package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Caller interface for AlphaVantage client
type Caller interface {
	GetLatestPriceForStock(symbol string) (*LatestStockPrice, error)
	GetStockInformation(symbol string) (*SearchStockAPIResponse, error)
}

// Client object for AlphaVantage configuration
type Client struct {
	EndpointURL string
	APIKey      string
}

// New creates a new client interface for AlphaVantage APIs
func New(apiKey string) *Client {
	var endpointURL = "https://www.alphavantage.co/query"

	return &Client{
		EndpointURL: endpointURL,
		APIKey:      apiKey,
	}
}

func (c *Client) GetStockInformation(symbol string) (*SearchStockAPIResponse, error) {
	response, err := http.Get(c.EndpointURL + "?function=SYMBOL_SEARCH&apikey=" + c.APIKey + "&keywords=" + symbol)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to fetch last price from API for symbol: %s", symbol)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read/parse response from API for symbol: %s", symbol)
	}

	// Parse the response from the API into a well defined type
	var searchSymbolMap map[string][]SearchStockAPIResponse
	if err := json.Unmarshal(responseData, &searchSymbolMap); err != nil {
		return nil, errors.Wrapf(err, "API limit exceeded: %s", symbol)
	}

	for _, stock := range searchSymbolMap["bestMatches"] {
		if stock.MatchScore == "1.0000" {
			matchedStock := stock
			return &matchedStock, nil
		}
	}

	return nil, errors.New("stock not found")
}

// GetLatestPriceForStock gets the latest price and volume information for symbol
func (c *Client) GetLatestPriceForStock(symbol string) (*LatestStockPrice, error) {
	response, err := http.Get(c.EndpointURL + "?function=GLOBAL_QUOTE&apikey=" + c.APIKey + "&symbol=" + symbol)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to fetch last price from API for symbol: %s", symbol)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read/parse response from API for symbol: %s", symbol)
	}

	// Parse the response from the API into a well defined type
	var latestPriceMap map[string]LatestPriceAPIResponse
	if err = json.Unmarshal(responseData, &latestPriceMap); err != nil {
		return nil, errors.Wrapf(err, "API limit exceeded: %s", symbol)
	}

	latestPriceSymbol := latestPriceMap["Global Quote"].Symbol
	// If API responded with an empty Global Quote field, it denotes that the symbol does not exist
	if symbol == "" {
		return nil, errors.New("invalid stock symbol")
	}

	stockInformation, err := c.GetStockInformation(latestPriceSymbol)
	if err != nil {
		return nil, errors.New("cannot get stock information")
	}

	// AlphaVantage API response for this endpoint returns a fixed Global Quote field on the body
	// This is to strip the relevant fields out of the Global Quote root
	return &LatestStockPrice{
		Symbol: latestPriceSymbol,
		Name:   stockInformation.Name,
		Price:  convertStringToFloat(latestPriceMap["Global Quote"].Open),
	}, nil
}

// convertStringToFloat converts string to float with a default to 0
func convertStringToFloat(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return num
}
