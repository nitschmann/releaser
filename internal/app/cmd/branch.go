package cmd

import (
	"fmt"

	cmdServ "github.com/nitschmann/releaser/internal/app/cmd/service"
	branchServ "github.com/nitschmann/releaser/internal/app/git/branch/service"
	"github.com/spf13/cobra"
)

func newBranchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "branch [TITLE]",
		Aliases: []string{"b"},
		Short:   "Creates a branch name with the given title and with configured path based specific rules",
		Long: `
Creates a branch name with the given title and with configured path based specific rules.
It could also be directly checked out.`,
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			checkout, err := cmd.Flags().GetBool("checkout")
			if err != nil {
				printCliErrorAndExit(err)
			}

			typeValue, err := cmd.Flags().GetString("type")
			if err != nil {
				printCliErrorAndExit(err)
			}

			templateValues := make(map[string]string)
			templateValues["Type"] = typeValue

			for _, flagName := range currentConfigRule.FlagNamesForBranch() {
				value, err := cmd.Flags().GetString(flagName)
				if err != nil {
					printCliErrorAndExit(err)
				}

				templateValues[flagName] = value
			}

			var title string
			if len(args) > 0 {
				title = args[0]
			}

			branchService := branchServ.New(currentConfigRule, GitService, templateValues)
			name, err := branchService.BuildNewName(typeValue, title)
			if err != nil {
				printCliErrorAndExit(err)
			}

			if checkout {
				err = branchService.Checkout(name)
				if err != nil {
					printCliErrorAndExit(err)
				}
			} else {
				fmt.Println(name)
			}
		},
	}

	cmd.Flags().BoolP("checkout", "c", false, "Set to true to checkout the new branch with generated name directly with git")
	cmd.Flags().StringP("type", "t", currentConfigRule.GitBranchDefaultType, fmt.Sprintf("Type of the branch which is one of %s", currentConfigRule.GetGitBranchTypes()))

	if currentConfigRule.GitBranchForceType {
		cmd.MarkFlagRequired("type")
	}

	if currentConfigRule.GitBranchDefaultType != "" {
		cmd.Flags().Set("type", currentConfigRule.GitBranchDefaultType)
	}

	cmdDyanamicFlagService, err := cmdServ.NewDynamicFlagConfig("branch", currentConfigRule)
	if err != nil {
		printCliErrorAndExit(err)
	}

	cmdDyanamicFlagService.AddFlagsForCmd(cmd)

	return cmd
}
