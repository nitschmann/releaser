package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/app/config"
	gitServ "github.com/nitschmann/releaser/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newFullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "full",
		Aliases: []string{"f", "full-log"},
		Short:   "Prints full release log including new version tag, changelog and compare URL",
		Long:    "Prints full release log including new version tag, changelog and compare URL",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(GitService, config.Get().FirstVersion)

			latestVersionTag, err := versionTagService.LatestVersionTag(config.Get().LatestVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			newVersionTag, err := versionTagService.CreateNew(config.Get().NewVersion)
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

			releaseService := gitServ.NewReleaseService(GitService, config.Get().GitRemote, config.Get().GitRepoURL)
			releaseTitle := releaseService.Title(newVersionTag)
			releaseCompareURL, err := releaseService.RepoVersionTagCompareURL(latestVersionTag, newVersionTag)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(releaseTitle + "\n")
			fmt.Printf("New version: %s\n", newVersionTag)

			if latestVersionTag != "" && releaseCompareURL != "" {
				fmt.Printf("Latest version: %s\n", latestVersionTag)
				fmt.Printf("Compare URL: %s\n", releaseCompareURL)
			}

			fmt.Println("\n## Changelog")
			fmt.Println("")
			for i := 0; i < len(changelog); i++ {
				fmt.Printf("* %s\n", changelog[i])
			}
		},
	}

	return cmd
}
