package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/nitschmann/releaser/internal/app/config"
	"github.com/spf13/cobra"
)

// DynamicFlagConfig is a service struct for handling dynamic flags for Cobra commands based on config.Rule
type DynamicFlagConfig struct {
	// The type of command for which the dynamic flags should be used for. Must be one of:
	// branch || commit
	CommandType string `validate:"oneof='branch' 'commit'"`
	configRule  config.Rule
}

// NewDynamicFlagConfig returns a new pointer instance of DynamicFlagConfig with the given arguments
func NewDynamicFlagConfig(commandType string, configRule config.Rule) (*DynamicFlagConfig, error) {
	d := &DynamicFlagConfig{
		CommandType: commandType,
		configRule:  configRule,
	}

	err := d.validate()
	if err != nil {
		return d, err
	}

	return d, nil
}

// AddFlagsForCmd adds the DynamicFlagConfig.configRule defined flags to the cmd
func (d DynamicFlagConfig) AddFlagsForCmd(cmd *cobra.Command) {
	var list []config.DynamicFlag

	if d.CommandType == "branch" {
		list = d.configRule.FlagsForBranch()
	} else if d.CommandType == "commit" {
		list = d.configRule.FlagsForCommit()
	}

	for _, flag := range list {
		cmd.Flags().String(flag.Name, flag.Default, flag.Description)

		if flag.Required {
			cmd.MarkFlagRequired(flag.Name)

			if flag.Default != "" {
				cmd.Flags().Set(flag.Name, flag.Default)
			}
		}
	}
}

func (d DynamicFlagConfig) validate() error {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		return err
	}

	return nil
}
