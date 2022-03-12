package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nitschmann/releaser/internal/apperror"
	configPkg "github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/pkg/git"
)

func getAutoYesyByFlag(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(yesFlagName)
}

func getCustomFlagValues(flags []configPkg.Flag, cmd *cobra.Command) (map[string]string, error) {
	var flagValues map[string]string = make(map[string]string)

	for _, flag := range flags {
		flagName := flag.GetName()
		flagValue, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return flagValues, err
		}

		flagValues[flagName] = flagValue
	}

	return flagValues, nil
}

func getCustomFlagValuesForBranch(cmd *cobra.Command) (map[string]string, error) {
	return getCustomFlagValues(config.GetFlagsForBranch(), cmd)
}

func getCustomFlagValuesForCommit(cmd *cobra.Command) (map[string]string, error) {
	return getCustomFlagValues(config.GetFlagsForCommit(), cmd)
}

func getCustomFlagValuesForRelease(cmd *cobra.Command) (map[string]string, error) {
	return getCustomFlagValues(config.GetFlagsForRelease(), cmd)
}

func getGitExecutableByFlag(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(gitExecutableFlagName)
}

func getTypeByFlag(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(typeFlagName)
}

func printCLIErrorAndExit(err error) {
	switch e := err.(type) {
	case *apperror.ConfigValidationErrors, *apperror.InvalidFlagValueError, *apperror.InvalidFlagError:
		fmt.Println(err.Error())
	case *apperror.PromptAbortError:
		fmt.Printf("%s\n", err.Error())
	case *git.CommandError:
		fmt.Printf("git execution error: \n%s\n", e.Output)
	default:
		fmt.Printf("An unexpected error occurred:\n%v\n", err.Error())
	}

	os.Exit(1)
}
