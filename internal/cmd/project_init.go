package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nitschmann/releaser/internal/service/project"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

func newProjectInitCmd() *cobra.Command {
	return cmdWithOptions(
		&cobra.Command{
			Use:     "init",
			Aliases: []string{"i"},
			Short:   "Initialize a new releaser project in the current directory",
			Long: `
Initialize a new releaser project in the current directory with a default config file under .releaser/
`,
			RunE: func(cmd *cobra.Command, args []string) error {
				ctx := cmd.Context()

				gitExecutable, err := getGitExecutableByFlag(cmd)
				if err != nil {
					return err
				}

				autoYes, err := getAutoYesyByFlag(cmd)
				if err != nil {
					return err
				}

				git := gitPkg.New(gitExecutable)
				service := project.NewInitService(git)

				return service.Call(ctx, autoYes)
			},
		},
		withGitExecutableFlag(),
		withPromptAutoYesFlag(),
	)
}
