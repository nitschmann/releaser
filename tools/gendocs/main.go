package main

import (
	"log"

	"github.com/nitschmann/releaser/internal/app/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	releaserCmd := cmd.NewRootCmd()
	releaserCmd.LoadSubCommands()
	err := doc.GenMarkdownTree(releaserCmd.Cmd, "docs/cli")
	if err != nil {
		log.Fatal(err)
	}
}
