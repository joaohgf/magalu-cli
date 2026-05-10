## Builds the project
build:
	@echo "Building the project..."
	@go mod tidy
	@go mod download
	@go mod vendor
	@go build -o ./cli ./cmd/tarefeiro

test:
	@echo "Running tests..."
	@go test -v ./... --tags=unit

test-output:
	@echo "Running tests with output..."
	@go test -v ./... --tags=unit -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html