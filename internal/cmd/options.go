package cmd

import "github.com/spf13/cobra"

type cmdOption func(*cobra.Command)

func cmdWithOptions(cmd *cobra.Command, opts ...cmdOption) *cobra.Command {
	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

func withGitExecutableFlag() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP("git-executable", "g", config.Git.GetExecutable(), "Git executable")
	}
}

func withPromptAutoYesFlag() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().BoolP("yes", "y", false, "Automatically answer prompt questions with 'yes'")
	}
}
