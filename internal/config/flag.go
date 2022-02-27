package config

import "fmt"

// Flag is a custom configured flag which could be used for various commands of releaser
type Flag struct {
	Name           string  `mapstructure:"name" yaml:"name" validate:"required,alphanum,lowercase"`
	Description    *string `mapstructure:"description" yaml:"description"`
	Required       *bool   `mapstructure:"required" yaml:"required"`
	SkipForBranch  *bool   `mapstructure:"skip_for_branch" yaml:"skip_for_branch"`
	SkipForCommit  *bool   `mapstructure:"skip_for_commit" yaml:"skip_for_commit"`
	SkipForRelease *bool   `mapstructure:"skip_for_release" yaml:"skip_for_release"`
}

// GetName returns the value of the Name field
func (f Flag) GetName() string {
	return f.Name
}

// GetDescription returns the value of the Description field if given, else default values
func (f Flag) GetDescription() string {
	if f.Description != nil {
		return *f.Description
	}

	return fmt.Sprintf("Custom flag '%s'", f.Name)
}

// GetRequired returns the value of the Required field if given, else default value
func (f Flag) GetRequired() bool {
	if f.Required != nil {
		return *f.Required
	}

	return FlagRequiredDefault
}

// GetSkipForBranch returns the value of the SkipForBranch field if given, else default value
func (f Flag) GetSkipForBranch() bool {
	if f.SkipForBranch != nil {
		return *f.SkipForBranch
	}

	return FlagSkipForBranchDefault
}

// GetSkipForCommit returns the value of the SkipForCommit field if given, else default value
func (f Flag) GetSkipForCommit() bool {
	if f.SkipForCommit != nil {
		return *f.SkipForCommit
	}

	return FlagSkipForCommitDefault
}

// GetkipForRelease returns the value of the SkipForRelease field if given, else default value
func (f Flag) GetSkipForRelease() bool {
	if f.SkipForRelease != nil {
		return *f.SkipForRelease
	}

	return FlagSkipForReleaseDefault
}
