package cmd

import (
	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage the releaser configuration",
		Long:  "Manage the configuration (file) of releaser",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newConfigFileCmd())
	cmd.AddCommand(newConfigShowCmd())

	return cmd
}
