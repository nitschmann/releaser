package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/app/config"
	gitServ "github.com/nitschmann/releaser/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newChangelogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "changelog",
		Aliases: []string{"release-changelog"},
		Short:   "Prints the changelog for the release",
		Long:    "Prints the changelog for the release (if present)",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(GitService, config.Get().FirstVersion)

			latestVersionTag, err := versionTagService.LatestVersionTag(config.Get().LatestVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			changelogService := gitServ.NewChangelogService(GitService, versionTagService)
			changelog, err := changelogService.ChangelogFromVersionTag(latestVersionTag)
			if err != nil {
				printCliErrorAndExit(err)
			}

			if len(changelog) == 0 {
				printCliErrorAndExit("No committed changes were found. Please ensure you are using the correct branch.")
			}

			fmt.Println("## Changelog")
			fmt.Println("")
			for i := 0; i < len(changelog); i++ {
				fmt.Printf("* %s\n", changelog[i])
			}
		},
	}

	return cmd
}
