# STS MFA CLI

<p align="center">
  <a href="https://github.com/browniebroke/stsmfa-cli/actions/workflows/ci.yml?query=branch%3Amain">
    <img src="https://img.shields.io/github/actions/workflow/status/browniebroke/stsmfa-cli/ci.yml?branch=main&label=CI&logo=github&style=flat-square" alt="CI Status" >
  </a>
  <a href="https://codecov.io/gh/browniebroke/stsmfa-cli">
    <img src="https://img.shields.io/codecov/c/github/browniebroke/stsmfa-cli.svg?logo=codecov&logoColor=fff&style=flat-square" alt="Test coverage percentage">
  </a>
</p>
<p align="center">
  <a href="https://python-poetry.org/">
    <img src="https://img.shields.io/endpoint?url=https://python-poetry.org/badge/v0.json" alt="Poetry">
  </a>
  <a href="https://github.com/astral-sh/ruff">
    <img src="https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/astral-sh/ruff/main/assets/badge/v2.json" alt="Ruff">
  </a>
  <a href="https://github.com/pre-commit/pre-commit">
    <img src="https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white&style=flat-square" alt="pre-commit">
  </a>
</p>
<p align="center">
  <a href="https://pypi.org/project/stsmfa-cli/">
    <img src="https://img.shields.io/pypi/v/stsmfa-cli.svg?logo=python&logoColor=fff&style=flat-square" alt="PyPI Version">
  </a>
  <img src="https://img.shields.io/pypi/pyversions/stsmfa-cli.svg?style=flat-square&logo=python&amp;logoColor=fff" alt="Supported Python versions">
  <img src="https://img.shields.io/pypi/l/stsmfa-cli.svg?style=flat-square" alt="License">
</p>

---

**Source Code**: <a href="https://github.com/browniebroke/stsmfa-cli" target="_blank">https://github.com/browniebroke/stsmfa-cli </a>

---

Creating temporary profiles for multi-factor auth (MFA) protected accounts using AWS STS is too hard. This is a small CLI that helps with that.

## Installation

Via Homebrew:

```bash
brew install browniebroke/tap/stsmfa-cli
```

Via pip, pipx, or your favourite Python package manager:

```bash
pip install stsmfa-cli
```

## Usage

The CLI is a simple command `stsmfa` that creates a profile for a temporary session protected by MFA.

Assuming your `~/.aws/credentials` file looks like this:

```ini
[my-profile-name]
aws_access_key_id = AKIAXXXXX
aws_secret_access_key = xxxx
mfa_serial = arn:aws:iam::123456789010:mfa/first.last
```

When running, for example:

```bash
stsmfa --profile my-profile-name 123456
```

This will create a session using the MFA serial defined under `my-profile-name` with the one-time password `123456`, and save the required AWS key, secret and token under as a new profile `my-profile-name-mfa` in you `~/.aws/credentials` file.

Now to use that session, you just need to set `AWS_PROFILE=my-profile-name-mfa`.

If your MFA serial is defined under the default profile, you don't need to specify the `--profile` option.

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

## Credits

This package was created with
[Copier](https://copier.readthedocs.io/) and the
[browniebroke/pypackage-template](https://github.com/browniebroke/pypackage-template)
project template.
