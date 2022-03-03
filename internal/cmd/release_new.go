package cmd

import "github.com/spf13/cobra"

func newReleaseNewCmd() *cobra.Command {
	cmd := cmdWithOptions(
		&cobra.Command{
			Use:     "new [OPTION]",
			Aliases: []string{"n"},
			Short:   "Generate new release",
			Long:    "Generate the data for a new release. Pass the option argument in case you want only parts of it.",
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()

				// Flags
				customFlags, err := getCustomFlagValuesForBranch(cmd)
				if err != nil {
					return err
				}

				gitExecutable, err := getGitExecutableByFlag(cmd)
				if err != nil {
					return err
				}

				return nil
			},
		},
		withCustomFlagsForRelease(),
		withGitExecutableFlag(),
	)

	cmd.Flags().String("first-tag", config.Release.GetFirstTag(), "First tag if none is present yet")
	cmd.Flags().String("git-remote", config.Git.GetRemote(), "Git remote")
	cmd.Flags().String("tag", "", "The new tag of the release")
	cmd.Flags().String("target", config.Release.GetTarget(), "Target (branch) of the release")

	return cmd
}
