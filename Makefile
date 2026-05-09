## Builds the project
build:
	@echo "Building the project..."
	@go mod tidy
	@go mod download
	@go mod vendor
	@go build -o ./cli ./cmd/tarefeiro

