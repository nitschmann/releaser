package cmd

import (
	"fmt"

	"github.com/nitschmann/releaser/internal/helper"
	"github.com/nitschmann/releaser/internal/service/release"
	"github.com/nitschmann/releaser/pkg/git"
	"github.com/spf13/cobra"
)

var releaseNewCmdWithArgs *cobra.Command = cmdWithOptions(
	&cobra.Command{
		Use:     "new [OPTION]",
		Aliases: []string{"n"},
		Args:    cobra.ExactArgs(1),
		Short:   "Generate new release",
		Long:    "Generate the data for a new release. Pass the option argument in case you want only parts of it.",
		RunE: func(cmd *cobra.Command, args []string) error {
			validOptions := []string{"description", "name", "tag", "target"}
			givenOption := args[0]

			if !helper.StringSliceIncludesElement(validOptions, givenOption) {
				return fmt.Errorf("Invalid option '%s' given. Valid options are %v", givenOption, validOptions)
			}

			ctx := cmd.Context()

			// Flags
			customFlags, err := getCustomFlagValuesForRelease(cmd)
			if err != nil {
				return err
			}

			firstTag, err := cmd.Flags().GetString("first-tag")
			if err != nil {
				return err
			}

			gitExecutable, err := getGitExecutableByFlag(cmd)
			if err != nil {
				return err
			}

			gitRemote, err := cmd.Flags().GetString("git-remote")
			if err != nil {
				return err
			}

			tag, err := cmd.Flags().GetString("tag")
			if err != nil {
				return err
			}

			target, err := cmd.Flags().GetString("target")
			if err != nil {
				return err
			}

			// service
			releaseService := release.NewGenerateService(git.New(gitExecutable))
			r, err := releaseService.Call(
				ctx,
				config,
				templateValues,
				customFlags,
				firstTag,
				gitRemote,
				tag,
				target,
			)
			if err != nil {
				return err
			}

			switch givenOption {
			case "description":
				fmt.Println(r.Description)
			case "name":
				fmt.Println(r.Name)
			case "tag":
				fmt.Println(r.Tag)
			case "target":
				fmt.Println(r.Target)
			}


			return nil
		},
	},
	withCustomFlagsForRelease(),
	withGitExecutableFlag(),
	withReleaseFlags(),
)
