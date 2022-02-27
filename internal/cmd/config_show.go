package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	configService "github.com/nitschmann/releaser/internal/service/config"
)

func newConfigShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show",
		Aliases: []string{"s"},
		Short:   "Show currently used configuration data structure",
		Long:    "Show the currently used configuration data structure for releaser in the current path",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			onlyYamlOption, err := cmd.Flags().GetBool("only-yaml")
			if err != nil {
				return err
			}

			showService := configService.NewShowService()
			configYamlData, err := showService.Call(ctx, config)
			if err != nil {
				return err
			}

			if !onlyYamlOption {
				var msg string
				if configFileUsed == "" {
					msg = fmt.Sprintf("No configuration file found which could be applied to '%s'. Usign the default:", currentCommandDir)
				} else {
					msg = fmt.Sprintf("Configuration file '%s' used under '%s':", configFileUsed, currentCommandDir)
				}

				fmt.Println(msg)
				fmt.Println("--- YAML ---")
			}

			fmt.Println(string(configYamlData))

			return nil
		},
	}

	cmd.Flags().Bool("only-yaml", false, "Only print the YAML output of the configuration")

	return cmd
}
