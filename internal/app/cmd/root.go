package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	cmdServ "github.com/nitschmann/release-log/internal/app/cmd/service"
	"github.com/nitschmann/release-log/internal/app/config"
	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// AppVersion is the global CLI application version
	AppVersion string
	// GitService for the cmd package
	GitService git.Git
	// Private vars
	cmdConfigRuleService *cmdServ.ConfigRule
	rootCmd              *cobra.Command

	_ = func() error {
		currentCmdDir, err := os.Getwd()
		if err != nil {
			printCliErrorAndExit(err)
		}

		initAppConfig()

		rootCmd = NewRootCmd()
		cmdConfigRuleService = cmdServ.NewConfigRule(currentCmdDir)

		config.SetDefaultValues()
		err = config.Load(false)
		if err != nil {
			printCliErrorAndExit(err)
		}

		return nil
	}()
)

// Execute is the app-wide CLI entrypoint
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printCliErrorAndExit(err)
	}
}

func initAppConfig() {
	godotenv.Load()
	config.Init()
}

// NewRootCmd returns the application and global facing root cobra command
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release-log",
		Short: "CLI tool for Git release version tags and logs",
		Long:  "CLI tool for Git release changelogs, logs and version tags",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := config.Load(true)
			if err != nil {
				printCliErrorAndExit(err)
			}

			GitService = git.New(config.Get().GitExecutable)
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().String("first-version", "v0.0.1", "The first release version which should be initially used")
	cmd.PersistentFlags().StringP("git-executable", "g", "git", "The system-wide used Git executable")
	cmd.PersistentFlags().StringP("git-remote", "r", "origin", "Git remote which should be used for comparison")
	cmd.PersistentFlags().StringP("git-repo-url", "u", "", "Git repository URL which could be overwritten. (If no URL is given the one of the git-remote is used)")
	cmd.PersistentFlags().String("new-version", "", "New Git release version tag to be used (if not given it will be detected automatically using git)")
	cmd.PersistentFlags().String("latest-version", "", "Latest Git release version tag to be used (if not given it will be detected automatically using git)")

	viper.BindPFlag("first_version", cmd.PersistentFlags().Lookup("first-version"))
	viper.BindPFlag("git_executable", cmd.PersistentFlags().Lookup("git-executable"))
	viper.BindPFlag("git_remote", cmd.PersistentFlags().Lookup("git-remote"))
	viper.BindPFlag("git_repo_url", cmd.PersistentFlags().Lookup("git-repo-url"))
	viper.BindPFlag("new_version", cmd.PersistentFlags().Lookup("new-version"))
	viper.BindPFlag("latest_version", cmd.PersistentFlags().Lookup("latest-version"))

	// Default static commands
	cmd.AddCommand(newChangelogCmd())
	cmd.AddCommand(newFullCmd())
	cmd.AddCommand(newLatestVersionCmd())
	cmd.AddCommand(newNewVersionCmd())
	cmd.AddCommand(newTitleCmd())
	cmd.AddCommand(newVersionCmd())

	return cmd
}

func printCliErrorAndExit(msg interface{}) {
	fmt.Println("An unexpected error occurred: ", msg)
	os.Exit(1)
}
