# STS MFA CLI (Go Version)

A Go rewrite of the Python STS MFA CLI tool. This is a command-line tool that creates temporary AWS profiles for multi-factor authentication (MFA) protected AWS accounts using AWS STS (Security Token Service).

## Features

- ✅ Same CLI API as the Python version
- ✅ Reads AWS credentials from `~/.aws/credentials`
- ✅ Validates MFA configuration
- ✅ Calls AWS STS with MFA token
- ✅ Writes temporary credentials to a new profile
- ✅ Colored terminal output
- ✅ Cross-platform builds
- ✅ Homebrew support

## Installation

### Via Homebrew (Recommended)

```bash
# Add the tap (once)
brew tap browniebroke/tap

# Install the tool
brew install stsmfa-cli
```

### Via Go Install

```bash
go install github.com/browniebroke/stsmfa-cli@latest
```

### Download Binary

Download the latest binary from the [releases page](https://github.com/browniebroke/stsmfa-cli/releases).

## Usage

The CLI API is identical to the Python version:

```bash
# Basic usage
stsmfa 123456

# Specify profile
stsmfa 123456 --profile my-profile

# Specify custom MFA profile name
stsmfa 123456 --profile my-profile --mfa-profile my-custom-mfa

# Alternative command name (for compatibility)
awsmfa 123456 --profile my-profile
```

### Options

- `--profile, -p`: The profile to use for obtaining the session token (default: "default")
- `--mfa-profile`: The profile to write the session data to (default: `<profile>-mfa`)
- `--version`: Show version information
- `--help`: Show help

## Requirements

- AWS credentials file at `~/.aws/credentials`
- MFA device configured in your AWS profile with `mfa_serial` parameter

### Example AWS Credentials File

```ini
[default]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY
mfa_serial = arn:aws:iam::123456789012:mfa/your-mfa-device

[my-profile]
aws_access_key_id = ANOTHER_ACCESS_KEY
aws_secret_access_key = ANOTHER_SECRET_KEY
mfa_serial = arn:aws:iam::123456789012:mfa/another-mfa-device
```

## Development

### Prerequisites

- Go 1.21 or later
- Make (optional, for build targets)

### Build

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install

# Install with symlinks for both command names
make install-links
```

### Testing

```bash
make test
```

### Dependencies

- [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2) - AWS API integration
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Color](https://github.com/fatih/color) - Terminal colors
- [go-ini](https://github.com/go-ini/ini) - INI file parsing

## Differences from Python Version

- Written in Go instead of Python
- Uses AWS SDK for Go v2 instead of boto3
- Uses Cobra instead of Typer for CLI
- Uses fatih/color instead of Rich for terminal output
- Single binary distribution instead of Python package
- Homebrew support out of the box
- Cross-platform binaries

## Migration from Python Version

The Go version is a drop-in replacement for the Python version. The CLI API is identical, so you can simply replace the Python version with the Go version without changing your usage patterns.

## License

MIT License - see LICENSE file for details.
