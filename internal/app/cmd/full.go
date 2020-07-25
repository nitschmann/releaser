package cmd

import (
	"fmt"

	"github.com/nitschmann/release-log/internal/app/config"
	gitServ "github.com/nitschmann/release-log/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newFullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "full",
		Aliases: []string{"f", "full-log"},
		Short:   "Prints full release log including new version tag, changelog and compare URL",
		Long:    "Prints full release log including new version tag, changelog and compare URL",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(config.Get().FirstVersion)

			latestVersioTag, err := versionTagService.LatestVersionTag(config.Get().LatestVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			_, err = versionTagService.BuildNew(config.Get().NewVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			logService := gitServ.NewLogService(versionTagService)
			changelog, err := logService.ChangelogFromVersionTag(latestVersioTag)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(changelog)
		},
	}

	return cmd
}
