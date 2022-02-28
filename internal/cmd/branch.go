package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/service/branch"
	"github.com/nitschmann/releaser/pkg/git"
	"github.com/spf13/cobra"
)

func newBranchCmd() *cobra.Command {
	cmd := cmdWithOptions(
		&cobra.Command{
			Use:     "branch [NAME]",
			Aliases: []string{"b"},
			Short:   "Generate branch name",
			Args:    cobra.ExactArgs(1),
			Long: `
Generate a new branch name based on the configuration and check it out (optional).
In case the branch name is invalid (based on git standards) it will raise a warning.
`,
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()

				// Get all inputs
				name := args[0]

				gitExecutable, err := getGitExecutableByFlag(cmd)
				if err != nil {
					return err
				}

				branchType, err := getTypeByFlag(cmd)
				if err != nil {
					return err
				}

				checkout, err := cmd.Flags().GetBool("checkout")
				if err != nil {
					return err
				}

				customFlags, err := getCustomFlagValuesForBranch(cmd)
				if err != nil {
					return err
				}

				// Service
				generateBranchService := branch.NewGenerateService(git.New(gitExecutable))
				newBranchName, err := generateBranchService.Call(
					ctx,
					config,
					templateValues,
					customFlags,
					name,
					branchType,
					checkout,
				)
				if err != nil {
					return err
				}

				if !checkout {
					fmt.Println(newBranchName)
				}

				return nil
			},
		},
		withCustomFlagsForBranch(),
		withGitExecutableFlag(),
		withTypeFlag(!config.Branch.GetAllowedWithoutType()),
	)

	cmd.Flags().BoolP("checkout", "c", false, "Checkout the new branch")

	return cmd
}
