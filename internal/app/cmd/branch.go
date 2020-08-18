package cmd

import (
	"fmt"

	cmdServ "github.com/nitschmann/release-log/internal/app/cmd/service"
	"github.com/nitschmann/release-log/internal/app/config"
	branchServ "github.com/nitschmann/release-log/internal/app/git/branch/service"
	"github.com/spf13/cobra"
)

func init() {
	err := config.Get().ValidateRules()
	if err != nil {
		printCliErrorAndExit(err)
	}

	rule, err := cmdConfigRuleService.CurrentRule()
	if err != nil {
		printCliErrorAndExit(err)
	}

	cmd := newBranchCmd(rule)

	cmdDyanamicFlagService, err := cmdServ.NewDynamicFlagConfig("branch", rule)
	if err != nil {
		printCliErrorAndExit(err)
	}

	cmdDyanamicFlagService.AddFlagsForCmd(cmd)

	rootCmd.AddCommand(cmd)
}

func newBranchCmd(rule config.Rule) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "branch [TITLE]",
		Aliases: []string{"b"},
		Short:   "Creates a branch name with the given title and with configured path based specific rules",
		Long:    "Creates a branch name with the given title and with configured path based specific rules",
		Args:    cobra.RangeArgs(0, 1),
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

			for _, flagName := range rule.FlagNamesForBranch() {
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

			branchService := branchServ.New(rule, GitService, templateValues)
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
	cmd.Flags().StringP("type", "t", rule.GitBranchDefaultType, fmt.Sprintf("Type of the branch which is one of %s", rule.GetGitBranchTypes()))

	if rule.GitBranchForceType {
		cmd.MarkFlagRequired("type")
	}

	if rule.GitBranchDefaultType != "" {
		cmd.Flags().Set("type", rule.GitBranchDefaultType)
	}

	return cmd
}
