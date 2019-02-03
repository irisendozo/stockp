# stockp-api

This is a stock portfolio API that can buy + sell stocks using real-time market data using AlphaVantage API.

## Technical Stack

- Go 1.11

## Development & Testing

### Requirements

- MacOSX machine for `modd` and instructions
- go v1.11

### Setting up your Go workspace

1. Install Go

```sh
brew install go
```

2. Add the Go bin directory on your path

```sh
export PATH=$PATH:$GOPATH/bin
```

### Run application

```sh
make run
```

### Run tests

```sh
make test
```

Note: A caveat of the tests

## Relevant Project Structure

- `server` contains all server bootstrap logic + routing to API endpoints
- `app` contains all the business logic for each API endpoint
  - `app/balance.go` handles all business logic for the balance domain. This consists of adding cash to balance,
    withdrawing and checking balance amount
  - `app/stocks.go` handles all business logic for the stocks domain. This consists of calling the Alphavantage API
    and getting real time market data for each stock
- `pkg/alphavantage` contains the interface layer for calling alphavantage API

## Core Business Logic

- Each transaction i.e. add balance + withdraw balance are "locally stored" as historical data. This means that no actual
  mutation of the previous state is happening on the stored. It is when the data (balance and owned stocks) are fetched that
  they are consolidated and aggregated to form relevant results.

## Known bugs & Improvements

- Error codes handling middleware
- Routing middleware so that app logic does not need to take care of writing HTTP responses
- Cleaning up structs and data model
- Cleaner way of handling API limit errors from Alphavantage
- Cleaner + more efficient way of iterating through the historical data
