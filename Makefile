help: ## Show available Make targets
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_.-]+:.*##/ {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: help build test test-output
build: ## Build the project binary
	@echo "Building the project..."
	@go mod tidy
	@go mod download
	@go mod vendor
	@go build -o ./ ./cmd/tarefeiro

test: ## Run unit tests
	@echo "Running tests..."
	@go test -v ./... --tags=unit

test-output: ## Run unit tests and generate HTML coverage report
	@echo "Running tests with output..."
	@go test -v ./... --tags=unit -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html