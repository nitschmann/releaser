package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"

	"github.com/nitschmann/releaser/internal/helper"
)

var (
	// Version is the global command line application version
	Version string

	rootCmd *RootCmd

	_ = func() error {
		rootCmd = NewRootCmd()
		rootCmd.InitSubCommands()

		return nil
	}()
)

// Execute runs the whole command application
func Execute() {
	currentCommandDir, err := os.Getwd()
	if err != nil {
		printCLIErrorAndExit(err)
	}

	ctx := context.Background()
	ctx = helper.NewContextWithCommandExecutionPath(ctx, currentCommandDir)

	err = rootCmd.Cmd.ExecuteContext(ctx)
	if err != nil {
		printCLIErrorAndExit(err)
	}
}

// RootCmd is the cmd package entry abstraction struct
type RootCmd struct {
	Cmd *cobra.Command
}

// NewRootCmd returns an new pointer instance of RootCmd with default values
func NewRootCmd() *RootCmd {
	cmd := &cobra.Command{
		Use:   "releaser",
		Short: "CLI tool for smart and rule based Git branch, commit and release log management",
		Long: `
releaser is CLI tool which allows you to manage Git branch and commit naming structures based on rule
configurations for paths. It helps to create and publish useful and well-managed releases with
their corresponding logs.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()
			}

			return nil
		},
	}

	return &RootCmd{cmd}
}

func (rootCmd *RootCmd) InitSubCommands() {
	cmd := rootCmd.Cmd
	// 'project' command
	cmd.AddCommand(newProjectCmd())
	// 'version' command
	cmd.AddCommand(newVersionCmd())
}
