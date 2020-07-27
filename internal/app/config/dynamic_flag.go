package config

import "github.com/go-playground/validator/v10"

// DynamicFlag holds a config for a dynamic flag which applies for a certain path.
// It should be used within a Rule.
type DynamicFlag struct {
	// Name of the itself
	Name string `validate:"required,alpha,lowercase"`
	// Default is an optional default value which should always be used for the flag
	Default string
	// Description of the flag itself (optional)
	Description string

	Required bool

	// SkipForBranch will not use this flag for any branch releated command and operation
	SkipForBranch bool `mapstructure:"skip_for_branch" yaml="skip_for_branch"`
	// SkipForCommit will not use this flag for any commit releated command and operation
	SkipForCommit bool `mapstructure:"skip_for_commit" yaml="skip_for_commit"`
}

// Validate runs the validators of the DynamicFlag instance
func (d DynamicFlag) Validate() error {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		return err
	}

	return nil
}
