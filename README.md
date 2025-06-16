# STS MFA CLI

<p align="center">
  <a href="https://github.com/browniebroke/stsmfa-cli/actions/workflows/test.yml?query=branch%3Amain">
    <img src="https://img.shields.io/github/actions/workflow/status/browniebroke/stsmfa-cli/test.yml?branch=main&label=Tests&logo=github&style=flat-square" alt="Test Status" >
  </a>
  <a href="https://github.com/browniebroke/stsmfa-cli/releases">
    <img src="https://img.shields.io/github/v/release/browniebroke/stsmfa-cli?logo=github&style=flat-square" alt="GitHub Release">
  </a>
</p>
<p align="center">
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go" alt="Go Version">
  </a>
  <a href="https://github.com/pre-commit/pre-commit">
    <img src="https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white&style=flat-square" alt="pre-commit">
  </a>
  <img src="https://img.shields.io/github/license/browniebroke/stsmfa-cli?style=flat-square" alt="License">
</p>

---

**Source Code**: <a href="https://github.com/browniebroke/stsmfa-cli" target="_blank">https://github.com/browniebroke/stsmfa-cli </a>

---

Creating temporary profiles for multi-factor auth (MFA) protected accounts using AWS STS is too hard. This is a small CLI written in Go that helps with that.

## Features

- ✅ Same CLI API as the original Python version
- ✅ Reads AWS credentials from `~/.aws/credentials`
- ✅ Validates MFA configuration
- ✅ Calls AWS STS with MFA token
- ✅ Writes temporary credentials to a new profile
- ✅ Colored terminal output
- ✅ Cross-platform builds (Linux, macOS, Windows, ARM64)
- ✅ Single binary distribution
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
go install github.com/browniebroke/stsmfa-cli/cmd/stsmfa@latest
go install github.com/browniebroke/stsmfa-cli/cmd/awsmfa@latest
```

### Download Binary

Download the latest binary from the [releases page](https://github.com/browniebroke/stsmfa-cli/releases).

## Usage

The CLI is a simple command `stsmfa` (or `awsmfa` for compatibility) that creates a profile for a temporary session protected by MFA.

Assuming your `~/.aws/credentials` file looks like this:

```ini
[my-profile-name]
aws_access_key_id = AKIAXXXXX
aws_secret_access_key = xxxx
mfa_serial = arn:aws:iam::123456789010:mfa/first.last
```

When running, for example:

```bash
# Using either command name
stsmfa 123456 --profile my-profile-name
awsmfa 123456 --profile my-profile-name
```

This will create a session using the MFA serial defined under `my-profile-name` with the one-time password `123456`, and save the required AWS key, secret and token as a new profile `my-profile-name-mfa` in your `~/.aws/credentials` file.

Now to use that session, you just need to set `AWS_PROFILE=my-profile-name-mfa`.

If your MFA serial is defined under the default profile, you don't need to specify the `--profile` option.

### Command Options

- `--profile, -p`: The profile to use for obtaining the session token (default: "default")
- `--mfa-profile`: The profile to write the session data to (default: `<profile>-mfa`)
- `--version`: Show version information
- `--help`: Show help

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

## Migration from Python Version

This Go version is a drop-in replacement for the original Python version. The CLI API is identical, so you can simply replace the Python version with the Go version without changing your usage patterns.

### Advantages over Python Version

- **Single binary** distribution (no Python runtime required)
- **Faster startup** (Go binary vs Python interpreter)
- **Smaller footprint** (statically linked binary)
- **Cross-platform builds** ready for distribution
- **Homebrew support** out of the box

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- prettier-ignore-start -->
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://browniebroke.com/"><img src="https://avatars.githubusercontent.com/u/861044?v=4?s=80" width="80px;" alt="Bruno Alla"/><br /><sub><b>Bruno Alla</b></sub></a><br /><a href="https://github.com/browniebroke/stsmfa-cli/commits?author=browniebroke" title="Code">💻</a> <a href="#ideas-browniebroke" title="Ideas, Planning, & Feedback">🤔</a> <a href="https://github.com/browniebroke/stsmfa-cli/commits?author=browniebroke" title="Documentation">📖</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
<!-- prettier-ignore-end -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

## License

MIT License - see LICENSE file for details.
