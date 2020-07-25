package main

import (
	"github.com/nitschmann/release-log/internal/app/cmd"
)

var Version string

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
