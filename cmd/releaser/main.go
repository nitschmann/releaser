package main

import (
	"github.com/nitschmann/releaser/internal/app/cmd"
)

// Version is the global version for the releaser CLI application
var Version string

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
