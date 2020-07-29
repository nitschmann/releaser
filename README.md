**ATTENTION**

PROJECT IS UNDER ACTIVE DEVELOPMENT AND NOT IN A STABLE STATE YET. THERE CAN BE ANYTIME BREAKING CHANGES!

# release-log

This is a small and simple CLI tool to handle Git release (change)logs and version tags. It uses native `git` commands wrapped within a Golang application. Various binaries for Mac and Linux are shipped as well.

This project was created as POC for a few things and is a result of those tryouts.

## Main features

* Simple commandline interface
* Works in every checked out repository
* Configuration globally (via config), per command call (via CLI flags) or via ENV variables possible.
* Print full release logs which could be used to create useful releases in GitHub or GitLab using commits
* Native `git` binding and overwriting (if custom or modfified git is used) possible
* Smart handling for version tags (automatically generated new version tags or initially if not defined yet).
* The tool does ***NEVER*** execute `fetch`, `commit`, `checkout`, `push` or any other 'potentially dangerous' `git` command - this is still up to you ;)
* The tool is really useful for CI environments and execution. 

## Requirements

* `git` (version `>= 2.0.0` recommended)
* Optional: `go` (version `>= 1.13.0` recommended)

## Installation

### Via automated script

This is the easiest way to install a system-wide executable `release-log` command. 

The [install script](scripts/install.sh) downloads the matching pre-compiled binary file from the [latest available release](https://github.com/nitschmann/release-log/releases) in this repository. It works for most Linux and Mac versions and architectures. 

#### Requirements

* `curl` is needed

#### Steps

1. Execute `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/nitschmann/release-log/master/scripts/install.sh)"`
2. The script places the `release-log` binary under `/usr/local/bin` - so it should be on most systems and users accessible

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

Usage documentation for the CLI could be found in a separate [folder](docs/cli/release-log.md).
