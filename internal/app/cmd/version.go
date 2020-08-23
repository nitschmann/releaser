package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Version of releaser",
		Long:    "The version of releaser itself",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(AppVersion)
		},
	}

	return cmd
}
