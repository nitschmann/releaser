package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/nitschmann/releaser/internal/service/release"
	"github.com/nitschmann/releaser/pkg/git"
	"github.com/spf13/cobra"
)

func newReleasePublishCmd() *cobra.Command {
	cmd := cmdWithOptions(
		&cobra.Command{
			Use:     "publish [UPSTREAM]",
			Aliases: []string{"p"},
			Args:    cobra.ExactArgs(1),
			Short:   "Publish a release",
			Long:    "Publish a release to a specific upstream.",
			RunE: func(cmd *cobra.Command, args []string) error {
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

				apiToken, err := cmd.Flags().GetString("api-token")
				if err != nil {
					return err
				}

				autoYes, err := getAutoYesyByFlag(cmd)
				if err != nil {
					return err
				}

				isDraft, err := cmd.Flags().GetBool("draft")
				if err != nil {
					return err
				}

				isPreRelease, err := cmd.Flags().GetBool("pre-release")
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
				releaseObj.IsDraft = isDraft
				releaseObj.IsPreRelease = isPreRelease

				result, err := release.NewPublishService().Call(
					ctx,
					config,
					releaseObj,
					args[0],
					apiToken,
					autoYes,
				)
				if err != nil {
					return err
				}

				resultJSON, err := json.Marshal(result)
				if err != nil {
					return err
				}

				fmt.Println(string(resultJSON))

				return nil
			},
		},
		withCustomFlagsForRelease(),
		withGitExecutableFlag(),
		withReleaseFlags(),
		withPromptAutoYesFlag(),
	)

	cmd.Flags().String("api-token", "", "API token for the choosen upstream")
	cmd.Flags().Bool("draft", false, "Release should be published as draft")
	cmd.Flags().Bool("pre-release", false, "Release should be published as pre-release")

	return cmd
}
