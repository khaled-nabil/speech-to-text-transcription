# Makefile
.PHONY: build clean deploy invoke-health lint deps

# Install dependencies
deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
lint:
	golangci-lint run ./...
