package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nitschmann/releaser/internal/service/project"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

func newProjectInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Short:   "Initialize a new releaser project in the current directory",
		Long: `
Initialize a new releaser project in the current directory with a default config file under .releaser/
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			gitExecutable, err := cmd.Flags().GetString("git-executable")
			if err != nil {
				return err
			}

			autoYes, err := cmd.Flags().GetBool("yes")
			if err != nil {
				return err
			}

			git := gitPkg.New(gitExecutable)
			service := project.NewInitService(git)

			return service.Call(ctx, autoYes)
		},
	}

	cmd.Flags().StringP("git-executable", "g", "git", "Git executable")
	cmd.Flags().BoolP("yes", "y", false, "Automatically answer prompt questions with 'yes'")

	return cmd
}
