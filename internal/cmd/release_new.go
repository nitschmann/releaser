package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/nitschmann/releaser/internal/service/release"
	"github.com/nitschmann/releaser/pkg/git"
	"github.com/spf13/cobra"
)

func newReleaseNewCmd() *cobra.Command {
	cmd := cmdWithOptions(
		&cobra.Command{
			Use:     "new [OPTION]",
			Aliases: []string{"n"},
			Args:    cobra.RangeArgs(0, 1),
			Short:   "Generate new release",
			Long:    "Generate the data for a new release. Pass the option argument in case you want only parts of it.",
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) > 0 {
					return releaseNewCmdWithArgs.RunE(cmd, args)
				} else {
					ctx := cmd.Context()

					// Flags
					customFlags, err := getCustomFlagValuesForRelease(cmd)
					if err != nil {
						return err
					}

					firstTag, err := cmd.Flags().GetString("first-tag")
					if err != nil {
						return err
					}

					gitExecutable, err := getGitExecutableByFlag(cmd)
					if err != nil {
						return err
					}

					gitRemote, err := cmd.Flags().GetString("git-remote")
					if err != nil {
						return err
					}

					tag, err := cmd.Flags().GetString("tag")
					if err != nil {
						return err
					}

					target, err := cmd.Flags().GetString("target")
					if err != nil {
						return err
					}

					jsonOnly, err := cmd.Flags().GetBool("json")
					if err != nil {
						return err
					}

					// service
					releaseService := release.NewGenerateService(git.New(gitExecutable))
					releaseObj, err := releaseService.Call(
						ctx,
						config,
						templateValues,
						customFlags,
						firstTag,
						gitRemote,
						tag,
						target,
					)

					if err != nil {
						return err
					}

					if jsonOnly {
						releaseJSON, err := json.Marshal(releaseObj)
						if err != nil {
							return err
						}

						fmt.Println(string(releaseJSON))
					} else {
						fmt.Println("Repository URL: " + releaseObj.RepoHttpURL)
						fmt.Println("Target: " + releaseObj.Target)
						fmt.Println("Tag: " + releaseObj.Tag)
						fmt.Println("Name: " + releaseObj.Name)
						fmt.Println("Description:")
						fmt.Println("\n" + releaseObj.Description)
					}

					return nil
				}
			},
		},
		withCustomFlagsForRelease(),
		withGitExecutableFlag(),
		withReleaseFlags(),
	)

	cmd.Flags().Bool("json", false, "Print the result as JSON")

	return cmd
}
