package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

func newReleaseLatestTagCmd() *cobra.Command {
	return cmdWithOptions(&cobra.Command{
		Use:     "latest-tag",
		Aliases: []string{"t"},
		Short:   "Print latest git tag",
		Long:    "Print the latest available git tag in this project",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitExecutable, err := getGitExecutableByFlag(cmd)
			if err != nil {
				return err
			}

			git := gitPkg.New(gitExecutable)
			gitTag := gitPkg.NewTag(git)
			latestTag, err := gitTag.Latest()
			if err != nil {
				return err
			}

			fmt.Println(latestTag)

			return nil
		},
	},
		withGitExecutableFlag(),
	)
}
