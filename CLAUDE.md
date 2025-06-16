# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**STS MFA CLI** is a Go command-line tool that creates temporary AWS profiles for multi-factor authentication (MFA) protected AWS accounts using AWS STS (Security Token Service). The tool reads AWS credentials from `~/.aws/credentials`, validates MFA configuration, calls AWS STS with the MFA token, and writes temporary credentials to a new profile.

## Development Commands

### Setup

```bash
# Install dependencies
go mod tidy

# Install pre-commit hooks for code quality
pre-commit install
```

### Testing

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...

# Run make targets
make test
```

### Code Quality

```bash
# Run all pre-commit hooks manually
pre-commit run -a

# Individual commands (handled by pre-commit):
# - go fmt for formatting
# - go vet for static analysis
# - go mod tidy for dependency management
```

### Building the CLI

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install

# Clean build artifacts
make clean
```

### Running the CLI

```bash
# After building
./bin/awsmfa 123456 --profile my-profile

# After local installation
awsmfa 123456 --profile my-profile
```

## Architecture & Code Structure

### Core Components

- **`cmd/awsmfa/main.go`**: Main awsmfa binary using Cobra framework
  - Single command handling the entire MFA workflow
  - AWS credentials file parsing with go-ini
  - AWS SDK for Go v2 STS client integration for session token generation
  - Colored terminal output using fatih/color

### Key Dependencies

- **Cobra**: Popular Go CLI framework for the command interface
- **AWS SDK for Go v2**: Official AWS SDK for STS API calls
- **fatih/color**: Terminal color formatting
- **go-ini**: INI file parsing for AWS credentials

### Testing Strategy

- **Go testing** package with standard test files
- Test coverage for core functionality
- Tests cover credentials file parsing and path validation
- Filesystem mocking using t.TempDir() for isolated tests

### Entry Points

The `awsmfa` binary is built from the main.go file in the cmd/awsmfa/ directory.

## Development Workflow

### Code Standards

- **Go 1.21+** minimum requirement
- **gofmt** for code formatting (enforced by pre-commit)
- **go vet** for static analysis
- **Conventional Commits** enforced via commitlint
- **Pre-commit hooks** run automatically on commit

### Release Process

- **GoReleaser** for automated cross-platform builds and releases
- **GitHub Actions** for CI/CD pipeline
- **Homebrew tap** integration for easy installation
- **Semantic versioning** based on git tags

### Key Configuration

- **go.mod**: Go module dependencies
- **.goreleaser.yml**: Release configuration for cross-platform builds
- **Makefile**: Build targets and development commands
- **.pre-commit-config.yaml**: Code quality hooks for Go projects
