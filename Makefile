# Makefile for go-blueprint CLI

BINARY_NAME=go-blueprint

all: build test lint

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go

test:
	@echo "Running tests..."
	@go test ./... -v

lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf response.json
	@rm -rf test-lint-project test-react-plain test-react-tailwind test-docs-project test-final

.PHONY: all build test lint clean
