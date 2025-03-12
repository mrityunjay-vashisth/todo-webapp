.PHONY: all generate build run clean test

# Default target
all: generate build

# Generate API code from OpenAPI specs
generate:
	@echo "Generating code from OpenAPI specs..."
	@go run gen.go

# Build the API server
build:
	@echo "Building API server..."
	@go build -o bin/server cmd/server/main.go

# Run the API server
run:
	@echo "Running API server on port 8080..."
	@go run cmd/server/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Help target
help:
	@echo "Available targets:"
	@echo "  make generate  - Generate API code from OpenAPI specs"
	@echo "  make build     - Build the API server"
	@echo "  make run       - Run the API server"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make test      - Run tests"