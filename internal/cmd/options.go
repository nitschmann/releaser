package cmd

import (
	"github.com/spf13/cobra"

	configPkg "github.com/nitschmann/releaser/internal/config"
)

var (
	gitExecutableFlagName string = "git-executable"
	typeFlagName          string = "type"
	yesFlagName           string = "yes"
)

type cmdOption func(*cobra.Command)

func assignFlagsToCmd(cmd *cobra.Command, flags []configPkg.Flag) {
	for _, flag := range flags {
		name := flag.GetName()
		defaultValue := flag.GetDefault()

		cmd.Flags().String(name, defaultValue, flag.GetDescription())

		if flag.GetRequired() {
			cmd.MarkFlagRequired(flag.GetName())
		}

		if defaultValue != "" {
			cmd.Flags().Set(name, defaultValue)
		}

	}
}

func cmdWithOptions(cmd *cobra.Command, opts ...cmdOption) *cobra.Command {
	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

func withCustomFlagsForBranch() cmdOption {
	return func(cmd *cobra.Command) {
		assignFlagsToCmd(cmd, config.GetFlagsForBranch())
	}
}

func withCustomFlagsForCommit() cmdOption {
	return func(cmd *cobra.Command) {
		assignFlagsToCmd(cmd, config.GetFlagsForCommit())
	}
}

func withGitExecutableFlag() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP(gitExecutableFlagName, "g", config.Git.GetExecutable(), "Git executable")
	}
}

func withPromptAutoYesFlag() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().BoolP(yesFlagName, "y", false, "Automatically answer prompt questions with 'yes'")
	}
}

func withTypeFlag(required bool) cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP(typeFlagName, "t", "", "Specify explicit type")

		if required {
			cmd.MarkFlagRequired(typeFlagName)
		}
	}
}
