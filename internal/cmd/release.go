package cmd

import "github.com/spf13/cobra"

func newReleaseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "release",
		Aliases: []string{"r"},
		Short:   "Manage releases",
		Long:    "Generate, manage and publish releases in the project",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newReleaseLatestTagCmd())
	cmd.AddCommand(newReleaseNewCmd())

	return cmd
}
