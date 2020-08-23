# releaser

Is a CLI tool that allows you to manage branch and commit naming structures based on certain configurations under paths. It helps to create and publish useful and well-managed releases with their corresponding logs.

It uses native `git` commands wrapped within a Golang application. Various binaries for Mac and Linux are shipped as well.

This project was created as POC for a few things and is a result of those tryouts.

**!! ATTENTION !!**

THE PROJECT IS UNDER ACTIVE DEVELOPMENT AND NOT IN A STABLE STATE YET. THERE CAN BE ANYTIME BREAKING CHANGES!

## Main features

* Simple commandline interface
* Works in every checked out repository
* Configuration globally (via config), per command call (via CLI flags) or via ENV variables possible.
* Path based configuration is possible.
* Print full release logs which could be used to create useful releases in GitHub or GitLab using commits
* Native `git` binding and overwriting (if custom or modfified git is used) possible
* Smart handling for version tags (automatically generated new version tags or initially if not defined yet).
* The tool is really useful for CI environments and execution.

## Requirements

* `git` (version `>= 2.0.0` recommended)
* Optional: `go` (version `>= 1.13.0` recommended)

## Installation

### Via automated script

This is the easiest way to install a system-wide executable `releaser` command.

The [install script](scripts/install.sh) downloads the matching pre-compiled binary file from the [latest available release](https://github.com/nitschmann/releaser/releases) in this repository. It works for most Linux and Mac versions and architectures.

#### Requirements

* `curl` is needed

#### Steps

1. Execute `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/nitschmann/releaser/master/scripts/install.sh)"`
2. The script places the `releaser` binary under `/usr/local/bin` - so it should be on most systems and users accessible

### Compile it

If you want to have some custom installation or don't trust the automated script you are able to compile the binary on your own. A powerful `Makefile` is in the repo included. Take a look for more details.

A few examples:


* `make build-latest` - to build all latest version for Linux and Mac
* `make build-latest-darwin` - to build all latest versions for Mac
* `make build-latest-linux` - to build all latest versions for Linux
* `make buld-latest-darwin-amd64` - build latest version for Mac with amd64 architecture

All compiled binaries are placed in the [.build folder](.build), including the target OS and architecture in the filename. Place them wherever it is useful for your needs with a proper name.

## Usage

### CLI

Usage documentation for the CLI could be found in a separate [folder](docs/cli/releaser.md).
