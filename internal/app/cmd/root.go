package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nitschmann/release-log/internal/app/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

// App-wide CLI entrypoint
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		printCliErrorAndExit(err)
	}
}

// Package initialization
func init() {
	cobra.OnInitialize(initAppConfig)
	rootCmd = newRootCmd()

	setRootCmdFlags()
	setRootCmdViperBindings()

	config.SetDefaultValues()
}

func initAppConfig() {
	godotenv.Load()
	config.Init()
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release-log",
		Short: "CLI tool for Git release versions and logs",
		Long:  "CLI tool for Git release changelogs, logs and versions",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := config.Load()
			if err != nil {
				printCliErrorAndExit(err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	cmd.AddCommand(newFullCmd())
	cmd.AddCommand(newLatestVersionCmd())
	cmd.AddCommand(newNewVersionCmd())
	cmd.AddCommand(newTtitleCmd())

	return cmd
}

func printCliErrorAndExit(msg interface{}) {
	fmt.Println("An unexpected error occurred: ", msg)
	os.Exit(1)
}

func setRootCmdFlags() {
	rootCmd.PersistentFlags().String("first-version", "v0.0.1", "The first release version which should be initally used")
	rootCmd.PersistentFlags().StringP("git-executable", "g", "git", "The system-wide used Git executable")
	rootCmd.PersistentFlags().StringP("git-remote", "r", "origin", "Git remote which should be used for comparison")
	rootCmd.PersistentFlags().StringP("git-repo-url", "u", "", "Git repository URL which could be overwritten. (If no URL is given the one of the git-remote is used)")
	rootCmd.PersistentFlags().String("new-version", "", "New Git release version tag to be used (if not given it will be detected automatically using git)")
	rootCmd.PersistentFlags().String("latest-version", "", "Latest Git release version tag to be used (if not given it will be detected automatically using git)")
}

func setRootCmdViperBindings() {
	viper.BindPFlag("first_version", rootCmd.PersistentFlags().Lookup("first-version"))
	viper.BindPFlag("git_executable", rootCmd.PersistentFlags().Lookup("git-executable"))
	viper.BindPFlag("git_remote", rootCmd.PersistentFlags().Lookup("git-remote"))
	viper.BindPFlag("git_repo_url", rootCmd.PersistentFlags().Lookup("git-repo-url"))
	viper.BindPFlag("new_version", rootCmd.PersistentFlags().Lookup("new-version"))
	viper.BindPFlag("latest_version", rootCmd.PersistentFlags().Lookup("latest-version"))
}
