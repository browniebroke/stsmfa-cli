.PHONY: build build-all clean test install

# Version can be overridden with make VERSION=x.y.z
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

# Default target
build:
	go build $(LDFLAGS) -o bin/awsmfa ./cmd/awsmfa

# Build for multiple platforms (useful for releases)
build-all: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-linux-amd64 ./cmd/awsmfa
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/awsmfa-linux-arm64 ./cmd/awsmfa
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-darwin-amd64 ./cmd/awsmfa
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/awsmfa-darwin-arm64 ./cmd/awsmfa
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-windows-amd64.exe ./cmd/awsmfa

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# Run tests and generate coverage report
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Install locally
install:
	go install $(LDFLAGS) ./cmd/awsmfa

# Run pre-commit hooks
pre-commit:
	pre-commit run --all-files
