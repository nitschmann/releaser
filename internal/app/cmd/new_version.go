package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/app/config"
	gitServ "github.com/nitschmann/releaser/internal/app/git/service"
	"github.com/spf13/cobra"
)

func newNewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new-version",
		Aliases: []string{"n", "new", "new-tag"},
		Short:   "Prints the new version tag",
		Long:    "Prints the new version tag",
		Run: func(cmd *cobra.Command, args []string) {
			versionTagService := gitServ.NewVersionTagService(GitService, config.Get().FirstVersion)
			newVersion, err := versionTagService.CreateNew(config.Get().NewVersion)
			if err != nil {
				printCliErrorAndExit(err)
			}

			fmt.Println(newVersion)
		},
	}

	return cmd
}
