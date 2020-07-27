package cmd

import (
	"fmt"

	"github.com/nitschmann/release-log/internal/app/config"
	gitServ "github.com/nitschmann/release-log/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newLatestVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "latest-version",
		Aliases: []string{"latest", "latest-tag"},
		Short:   "Prints the latest available version tag",
		Long:    "Prints the latest available version tag. Is empty if no tag is available yet.",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(config.Get().FirstVersion)
			latestVersion, err := versionTagService.LatestVersionTag(config.Get().LatestVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(latestVersion)
		},
	}

	return cmd
}
