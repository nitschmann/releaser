package cmd

import "github.com/spf13/cobra"

type releaseCmd struct {
	cmd *cobra.Command
}

func newReleaseCmd() *releaseCmd {
	cmd := &cobra.Command{
		Use:     "release",
		Aliases: []string{"r"},
	}

	return &releaseCmd{cmd: cmd}
}
