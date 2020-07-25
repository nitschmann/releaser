package main

import (
	"log"

	"github.com/nitschmann/release-log/internal/app/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	releaseLogCmd := cmd.NewRootCmd()
	err := doc.GenMarkdownTree(releaseLogCmd, "docs/cli")
	if err != nil {
		log.Fatal(err)
	}
}
