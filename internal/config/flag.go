package config

// Flag is a custom configured flag which could be used fpr various commands of releaser
type Flag struct {
	Name           string  `mapstructure:"name" yaml:"name" validate:"required,alphanum,lowercase"`
	Description    *string `mapstructure:"description" yaml:"description"`
	Required       *bool   `mapstructure:"required" yaml:"required"`
	SkipForBranch  *bool   `mapstructure:"skip_for_branch" yaml:"skip_for_branch"`
	SkipForCommit  *bool   `mapstructure:"skip_for_commit" yaml:"skip_for_commit"`
	SkipForRelease *bool   `mapstructure:"skip_for_release" yaml:"skip_for_release"`
}
