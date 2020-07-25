package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Version of release-log",
		Long:    "The version of release-log itself",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(AppVersion)
		},
	}

	return cmd
}
