#!/bin/bash

# This script builds the Go application into a binary for the given (via input parameter) OS and arch

# --- Helper methods
contains_element () {
  local word=$1
  shift
  for e in "$@"; do [[ "$e" == "$word" ]] && return 0; done
  return 1
}

# --- Support lists
supported_command=( latest new-version )
supported_os=( darwin linux )
supported_architecture=( 386 amd64 arm arm64 )

# --- User inputs
command=$1
given_os=$2
given_arch=$3
version=$4

if ! contains_element ${command} "${supported_command[@]}"; then
  eco "Command is not supported!"
  exit 1
fi

if ! contains_element ${given_os} "${supported_os[@]}"; then
  echo "Given OS is not supported!"
  exit 1
fi


if ! contains_element ${given_arch} "${supported_architecture[@]}"; then
  echo "Given arch is not supported!"
  exit 1
fi

if [ "$command" == "latest" ] && [ -z "$version" ]; then
  version=$(go run cmd/releaselog/main.go latest-version)
  version="${version:-v0.0.1}"
fi

if [ "$command" == "new-version" ] && [ -z "$version" ]; then
  version=$(go run cmd/releaselog/main.go new-version)
fi

executable_name="release-log-$given_os-$given_arch"
output_path="./.build/$executable_name"
package_path="./cmd/releaselog/main.go"

env CGO_ENABLED=0 \
  GOOS=$given_os \
  GOARCH=$given_arch \
  go build \
  -o $output_path \
  -v -ldflags="-X 'main.Version=$version'" $package_path

if [ $? -ne 0 ]; then
  echo "An error occurred. Could not finish build process."
  exit 1
else
  printf "=== Finished: $output_path \n"
fi

