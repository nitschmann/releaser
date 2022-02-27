package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newConfigFileCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "file",
		Aliases: []string{"f"},
		Short:   "Print currently used configuration filepath",
		Long:    "Print currently used releaser configuration filepath",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(configFileUsed)
		},
	}
}
