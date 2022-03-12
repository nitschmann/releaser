package config

import (
	"reflect"
	"strings"

	validatorPkg "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/internal/validation"
)

var (
	// ConfigFileLookupPaths define the filepaths where to look for the releaser config file
	ConfigFileLookupPaths []string = []string{
		"./.releaser",
		"$HOME/.releaser",
		"~/.releaser",
		"/etc/releaser",
	}

	// Branch config default values
	BranchAllowedWithoutTypeDefault bool     = true
	BranchDelimiterDefault          string   = "-"
	BranchNameFormatDefault         string   = "{{if .Type}}{{ .Type }}{{end}} {{ .BranchName }}"
	BranchTypesDefault              []string = []string{"bug", "feature", "fix", "hotfix"}
	// Commit config default values
	CommitAllowedWithoutTypeDefault bool     = true
	CommitMessageFormatDefault      string   = "{{if .Type}}{{ .Type | ToTitle }}:{{end}} {{ .CommitMessage }}"
	CommitTypesDefault              []string = []string{"adjustment", "bug", "feature", "fix", "hotfix"}
	// Flag config default values
	FlagRequiredDefault       bool = false
	FlagSkipForBranchDefault  bool = false
	FlagSkipForCommitDefault  bool = false
	FlagSkipForReleaseDefault bool = true
	// Git config default values
	GitExecutableDefault string = "git"
	GitRemoteDefault     string = "origin"
	// Release config default values
	ReleaseFirstTagDefault          string = "v0.0.1"
	ReleaseNameFormatDefault        string = "Release {{ .ReleaseTag }} ({{ .DTYear }}-{{ .DTMonth }}-{{ .DTDay }})"
	ReleaseDescriptionFormatDefault string = `
## Changelog

{{range .GitCommitLogs -}}
* {{ .Message }}
{{end}}`
	ReleaseTargetDefault string = "master"
)

// Config has all the relevant settings
type Config struct {
	// Git specific config fields
	Git Git `mapstructure:"git" yaml:"git" validate:"required,dive"`
	// Branch specific config fields
	Branch Branch `mapstructure:"branch" yaml:"branch" validate:"required,dive"`
	// Commit specific config fields
	Commit Commit `mapstructure:"commit" yaml:"commit" validate:"required,dive"`
	// Release specific config fields
	Release Release `mapstructure:"release" yaml:"release" validate:"required,dive"`
	// Flags specify custom flags for commands
	Flags []Flag `mapstructure:"flags" yaml:"flags" validate:"dive"`
}

// New returns an new instance of Config with default values
func New() Config {
	return Config{
		Git:     newGit(),
		Branch:  newBranch(),
		Commit:  newCommit(),
		Release: newRelease(),
		Flags:   []Flag{},
	}
}

// Init loads the config from a file if found and returns the used config filepath as result
func Init() (string, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	for _, path := range ConfigFileLookupPaths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return "", err
		}
	}

	return viper.ConfigFileUsed(), nil
}

// GetFlagsForBranch returns all entries of the Flags field where SkipForBranch is not true
func (c Config) GetFlagsForBranch() FlagList {
	var flags FlagList

	for _, flag := range c.Flags {
		if !flag.GetSkipForBranch() {
			flags = append(flags, flag)
		}
	}

	return flags
}

// GetFlagsForCommit returns all entries of the Flags field where SkipForCommit is not true
func (c Config) GetFlagsForCommit() FlagList {
	var flags FlagList

	for _, flag := range c.Flags {
		if !flag.GetSkipForCommit() {
			flags = append(flags, flag)
		}
	}

	return flags
}

// GetFlagsForRelease returns all entries of the Flags field where SkipForRelease is not true
func (c Config) GetFlagsForRelease() FlagList {
	var flags FlagList

	for _, flag := range c.Flags {
		if !flag.GetSkipForRelease() {
			flags = append(flags, flag)
		}
	}

	return flags
}

// Validate the Config structure
func (c Config) Validate() error {
	validationTranslator, err := validation.NewTranslatorEN()
	if err != nil {
		return err
	}

	validator, err := validation.NewValidator(validationTranslator)
	if err != nil {
		return err
	}
	// Use the YAML tag as field for this validation
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		return name
	})

	err = validator.Struct(c)
	if err != nil {
		errs := err.(validatorPkg.ValidationErrors)
		validationErrors := apperror.NewConfigValidationErrors()

		for _, e := range errs {
			errMsg := e.Translate(validationTranslator)
			validationErrors.Add(apperror.NewConfigValidationError(e, errMsg))
		}

		return validationErrors
	}

	return nil
}
