.PHONY: build build-all clean test install

# Version can be overridden with make VERSION=x.y.z
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

# Default target
build:
	go build $(LDFLAGS) -o bin/stsmfa ./cmd/stsmfa
	go build $(LDFLAGS) -o bin/awsmfa ./cmd/awsmfa

# Build for multiple platforms (useful for releases)
build-all: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/stsmfa-linux-amd64 ./cmd/stsmfa
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-linux-amd64 ./cmd/awsmfa
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/stsmfa-linux-arm64 ./cmd/stsmfa
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/awsmfa-linux-arm64 ./cmd/awsmfa
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/stsmfa-darwin-amd64 ./cmd/stsmfa
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-darwin-amd64 ./cmd/awsmfa
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/stsmfa-darwin-arm64 ./cmd/stsmfa
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/awsmfa-darwin-arm64 ./cmd/awsmfa
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/stsmfa-windows-amd64.exe ./cmd/stsmfa
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/awsmfa-windows-amd64.exe ./cmd/awsmfa

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test -v ./...

# Install locally
install:
	go install $(LDFLAGS) ./cmd/stsmfa
	go install $(LDFLAGS) ./cmd/awsmfa
