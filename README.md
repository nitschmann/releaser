# releaser

A CLI to standardize Git branches, commit messages, and release notes across projects. It generates consistent names from templates, validates inputs, and can publish releases to upstreams like GitHub/GitLab.

## Features

- Generate and optionally checkout branches with safe names
- Create commit messages from templates and flags
- Generate release name/description from commit logs and templates
- Publish releases to upstream providers (configurable)
- Project-level overrides via path-based configuration

## Requirements

- Go 1.25+
- Git installed and available in PATH

## Install

### Option 1: Prebuilt binaries (Releases)

Download the latest binary for your platform from the [Releases](https://github.com/nitschmann/releaser/releases/latest) page.

Example (macOS arm64):

```bash
curl -L -o releaser https://github.com/nitschmann/releaser/releases/latest/download/releaser-darwin-arm64
chmod +x releaser
sudo mv releaser /usr/local/bin/
releaser version
```

Example (Linux amd64):

```bash
curl -L -o releaser https://github.com/nitschmann/releaser/releases/latest/download/releaser-linux-amd64
chmod +x releaser
sudo mv releaser /usr/local/bin/
releaser version
```

### Option 2: Install script (fetches latest release)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/nitschmann/releaser/master/scripts/install.sh)"
releaser version
```

### Option 3: Build from source

Build locally and place the binary on your PATH:

```bash
make build            # local builds for darwin + linux
# or
make build-latest     # builds using latest tag resolution

# binaries are written to ./.build/
```

Alternatively, run tests and lints:

```bash
make run-tests
make lint             # or: make lint-docker
```

## Usage (CLI)
See the generated command docs in [docs/cli/](docs/cli/) (entry file [docs/cli/README.md](docs/cli/README.md)).
Common commands:

```bash
# Initialize a project config in the current git repo
releaser project init

# Generate a branch name (prints) or checkout with -c
releaser branch "Feature X" -t fix
releaser branch -c "API v2" --type feature

# Generate a commit message
releaser commit "Add metrics" -t fix

# Generate release data (name/description) and print parts
releaser release new name
releaser release new description

# Publish a release
releaser release publish
```

## Configuration
releaser reads YAML from the following locations (first found wins):
- `./.releaser/config.yaml`
- `$HOME/.releaser/config.yaml`
- `/etc/releaser/config.yaml`

The configuration is flexible and supports project-specific overrides by path. Because it is fairly detailed, the structure and examples are documented separately.

- See [docs/CONFIGURATION.md](docs/CONFIGURATION.md) for a full reference and examples.

## Contributing

- Run lints and tests locally before opening PRs
- Keep templates and defaults backward compatible where possible

## License

Licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
