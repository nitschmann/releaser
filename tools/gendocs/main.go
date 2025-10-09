package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cobraDoc "github.com/spf13/cobra/doc"

	cmdPkg "github.com/nitschmann/releaser/internal/cmd"
)

func main() {
	outDir := filepath.Join("docs", "cli")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create docs directory: %v\n", err)
		os.Exit(1)
	}

	root := cmdPkg.NewRootCmd()
	root.InitSubCommands()
	root.Cmd.DisableAutoGenTag = true

	filePrepender := func(filename string) string { return "" }
	linkHandler := func(link string) string {
		if strings.HasSuffix(link, "releaser.md") {
			return strings.TrimSuffix(link, "releaser.md") + "README.md"
		}
		return link
	}

	if err := cobraDoc.GenMarkdownTreeCustom(root.Cmd, outDir, filePrepender, linkHandler); err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate cobra docs: %v\n", err)
		os.Exit(1)
	}

	_ = os.Rename(filepath.Join(outDir, "releaser.md"), filepath.Join(outDir, "README.md"))
}
