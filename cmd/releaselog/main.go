package main

import (
	"github.com/nitschmann/release-log/internal/app/cmd"
)

// Version is the global version for the release-log CLI application
var Version string

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
