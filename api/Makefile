.PHONY: run
run: ## Runs the service in development and watches for changes
ifeq ($(shell which modd),)
	brew install modd
endif
	modd

.PHONY: lint
lint: ## Runs more than 20 different linters using golangci-lint to ensure consistency in code.
ifeq ($(shell which golangci-lint),)
	brew install golangci/tap/golangci-lint
endif
	golangci-lint run

.PHONY: test
test: lint## Runs the tests and higlights race conditions
	GO111MODULE=on go test -race $$(go list ./...)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
