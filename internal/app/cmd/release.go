package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type releaseCmd struct {
	cmd *cobra.Command
}

func newReleaseCmd() *releaseCmd {
	cmd := &cobra.Command{
		Use:     "release",
		Aliases: []string{"r"},
		Short:   "Release management and logs",
		Long: `
Manages the releases and their corresponding logs and versions. Check the sub commands for detailed
descriptions.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	return &releaseCmd{cmd: cmd}
}

func (r *releaseCmd) loadSubCommands() {
}
