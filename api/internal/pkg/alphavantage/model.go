package alphavantage

// LatestPriceAPIResponse defines the API response structure from AlphaVantage's latest price endpoint
// https://www.alphavantage.co/documentation/#latestprice
type LatestPriceAPIResponse struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

// LatestStockPrice defines the latest price of a stock
type LatestStockPrice struct {
	Symbol string
	Name   string
	Price  float64
}

// SearchStockAPIResponse defines the API response structure from AlphaVantage's search endpoint
// https://www.alphavantage.co/documentation/#symbolsearch
type SearchStockAPIResponse struct {
	Symbol      string `json:"1. symbol"`
	Name        string `json:"2. name"`
	Type        string `json:"3. type"`
	Region      string `json:"4. region"`
	MarketOpen  string `json:"5. marketOpen"`
	MarketClose string `json:"6. marketClose"`
	Timezone    string `json:"7. timezone"`
	Currency    string `json:"8. currency"`
	MatchScore  string `json:"9. matchScore"`
}
