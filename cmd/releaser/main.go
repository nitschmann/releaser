package main

import (
	"github.com/nitschmann/releaser/internal/cmd"
)

// Version is the global version for the releaser CLI application
var Version string

func main() {
	cmd.Version = Version
	cmd.Execute()
}
