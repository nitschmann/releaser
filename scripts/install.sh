#!/bin/bash

# This is the installation script for the release-log binary. It downloads it automatically from the
# latest release available within the repo.

# --- Helper methods
contains_element () {
  local word=$1
  shift
  for e in "$@"; do [[ "$e" == "$word" ]] && return 0; done
  return 1
}

# --- Support lists
supported_os=( darwin linux )
supported_architecture=( 386 amd64 arm arm64 )
releases_base_url="https://github.com/nitschmann/release-log/releases"
latest_release_url="$releases_base_url/latest"

case $(uname -s) in
  Linux*)   machine_os="linux";;
  Darwin*)  machine_os="darwin";;
  *)        machine_os="UNKNOWN:${uname_machine_out}"
esac

if ! contains_element ${machine_os} "${supported_os[@]}"; then
  echo "The OS of this machine is not supported!"
  exit 1
fi

architecture=""
case $(uname -m) in
    i386)   architecture="386";;
    i686)   architecture="386";;
    x86_64) architecture="amd64";;
    arm)    dpkg --print-architecture | grep -q "arm64" && architecture="arm64" || architecture="arm";;
esac

if ! contains_element ${architecture} "${supported_architecture[@]}"; then
  echo "The architecture of this machine is not supported!"
  exit 1
fi

latest_release_tag_url=$(curl -Ls -w %{url_effective} -o /dev/null $latest_release_url)
latest_release_tag=${latest_release_tag_url##*/}
binary_name="release-log-$machine_os-$architecture"
binary_download_url="$releases_base_url/download/$latest_release_tag/$binary_name"

curl -o /usr/local/bin/release-log -L $binary_download_url
chmod +x /usr/local/bin/release-log
