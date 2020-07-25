package cmd

import (
	"fmt"

	"github.com/nitschmann/release-log/internal/app/config"
	gitServ "github.com/nitschmann/release-log/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newTtitleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "title",
		Aliases: []string{"t", "release-title"},
		Short:   "Prints the title of the new release",
		Long:    "Prints the title of the new release",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(config.Get().FirstVersion)
			newVersionTag, err := versionTagService.CreateNew(config.Get().NewVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			releaseService := gitServ.NewReleaseService(config.Get().GitRemote, config.Get().GitRepoUrl)
			releaseTitle := releaseService.Title(newVersionTag)

			fmt.Println(releaseTitle)
		},
	}

	return cmd
}
