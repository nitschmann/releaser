package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/service/commit"
	"github.com/nitschmann/releaser/pkg/git"
	"github.com/spf13/cobra"
)

func newCommitCmd() *cobra.Command {
	cmd := cmdWithOptions(
		&cobra.Command{
			Use:     "commit [MESSAGE]",
			Aliases: []string{"c"},
			Args:    cobra.ExactArgs(1),
			Short:   "Create new commit",
			Long:    "Create a new commit with the given message based on the current configuration",
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()

				// Get all inputs
				message := args[0]

				gitExecutable, err := getGitExecutableByFlag(cmd)
				if err != nil {
					return err
				}

				commitType, err := getTypeByFlag(cmd)
				if err != nil {
					return err
				}

				customFlags, err := getCustomFlagValuesForCommit(cmd)
				if err != nil {
					return err
				}

				onlyMessage, err := cmd.Flags().GetBool("message")
				if err != nil {
					return err
				}

				// Service
				createCommitService := commit.NewCreateService(git.New(gitExecutable))
				commitMessage, err := createCommitService.Call(
					ctx,
					config,
					templateValues,
					customFlags,
					message,
					commitType,
					onlyMessage,
				)
				if err != nil {
					return err
				}

				fmt.Println(commitMessage)

				return nil
			},
		},
		withCustomFlagsForCommit(),
		withGitExecutableFlag(),
		withTypeFlag(!config.Commit.GetAllowedWithoutType()),
	)

	cmd.Flags().BoolP("message", "m", false, "Print only the message of the new commit")

	return cmd
}
