package cmd

import (
	"github.com/spf13/cobra"
)

func newProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"p"},
		Short:   "Manage the releaser project",
		Long: `
Manage your releaser project.
Usually all the sub-commands should be executed within an initialized git directory and repository.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newProjectInitCmd())

	return cmd
}
