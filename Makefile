.PHONY: build run test clean coverage help

# Default target
all: build

# Build the application
build:
	@echo "Building tax-calculation API..."
	go build -o tax-api main.go

# Run the application
run:
	@echo "Running tax-calculation API..."
	go run main.go

# Run all tests
test:
	@echo "Running tests..."
	go test ./... -v

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f tax-api
	rm -f coverage.out coverage.html
	go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t tax-calculation-api:latest .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 tax-calculation-api:latest

# Help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  test          - Run all tests"
	@echo "  coverage      - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Install dependencies"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  help          - Show this help message"
