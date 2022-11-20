GOCMD=go
GOBUILD=$(GOCMD) build -race
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_ARTIFACTS = ".build-artifacts"
BINARY_NAME = "sbmf"

setup: ## Install tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0

lint: ## Run the linters
	golangci-lint run

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done


clean: ## cleans the build folder
	rm -f $(BUILD_ARTIFACTS)/$(BINARY_NAME)
	 
# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


.DEFAULT_GOAL := help