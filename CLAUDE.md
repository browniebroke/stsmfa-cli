# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**STS MFA CLI** is a Python command-line tool that creates temporary AWS profiles for multi-factor authentication (MFA) protected AWS accounts using AWS STS (Security Token Service). The tool reads AWS credentials from `~/.aws/credentials`, validates MFA configuration, calls AWS STS with the MFA token, and writes temporary credentials to a new profile.

## Development Commands

### Setup

```bash
# Install dependencies and set up development environment
uv sync

# Install pre-commit hooks for code quality
pre-commit install
```

### Testing

```bash
# Run all tests with coverage
uv run pytest

# Run specific test file
uv run pytest tests/test_cli.py

# Run tests with verbose output
uv run pytest -v
```

### Code Quality

```bash
# Run all pre-commit hooks manually
pre-commit run -a

# Individual linting commands (handled by pre-commit):
# - Ruff for linting and formatting (configured in pyproject.toml)
# - MyPy for type checking
# - Commitlint for conventional commit messages
```

### Running the CLI

```bash
# Install in development mode
uv sync

# Run via module
python -m stsmfa --help

# Or directly via entry points after installation:
stsmfa 123456 --profile my-profile
awsmfa 123456 --profile my-profile  # Alternative command name
```

## Architecture & Code Structure

### Core Components

- **`src/stsmfa/cli.py`**: Main CLI logic using Typer framework
  - Single command function `run()` that handles the entire MFA workflow
  - AWS credentials file parsing with configparser
  - Boto3 STS client integration for session token generation
  - Rich terminal output for user feedback

### Key Dependencies

- **Typer**: Modern CLI framework for the command interface
- **Boto3**: AWS SDK for STS API calls
- **Rich**: Terminal formatting and colored output
- **configparser**: AWS credentials file manipulation

### Testing Strategy

- **pytest** with comprehensive mocking using pytest-mock
- **pyfakefs** for filesystem mocking in tests
- Test coverage configured with 100% requirement (excluding `__main__.py`)
- Tests cover all error conditions and success paths

### Entry Points

Both `stsmfa` and `awsmfa` commands point to `stsmfa.cli:app` (defined in pyproject.toml scripts section).

## Development Workflow

### Code Standards

- **Python 3.9+** minimum requirement
- **Ruff** for all linting and formatting (replaces flake8, isort, black)
- **MyPy** with strict type checking enabled
- **Conventional Commits** enforced via commitlint
- **Pre-commit hooks** run automatically on commit

### Release Process

- **Semantic Release** with automated versioning based on conventional commits
- Version updates in both `pyproject.toml` and `src/stsmfa/__init__.py`
- Automated PyPI publishing via GitHub Actions

### Key Configuration

- All tool configuration centralized in `pyproject.toml`
- Coverage threshold enforced in CI
- Type checking with strict settings for production code, relaxed for tests
