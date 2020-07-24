package cmd

import (
	"fmt"

	"github.com.com/nitschmann/release-log/internal/app/config"
	gitServ "github.com.com/nitschmann/release-log/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newFullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "full",
		Aliases: []string{"f"},
		Short:   "Prints full release log including new version tag, changelog and compare URL",
		Long:    "Prints full release log including new version tag, changelog and compare URL",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(config.Get().FirstVersion)
			newVersion, err := versionTagService.BuildNew()
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(newVersion)
		},
	}

	return cmd
}
