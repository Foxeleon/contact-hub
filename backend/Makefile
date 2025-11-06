# Makefile for the backend service

.DEFAULT_GOAL := help

# Command to run the development server
run:
	@echo "Starting development server..."
	go run ./cmd/server/main.go

# Command to run all tests
test:
	@echo "Running all tests..."
	go test ./... -v

# Command to tidy up go.mod and go.sum
tidy:
	@echo "Tidying module dependencies..."
	go mod tidy

# Display help message
help:
	@echo "Available commands:"
	@echo "  make run    - Start the development server"
	@echo "  make test   - Run all unit tests"
	@echo "  make tidy   - Tidy module dependencies"