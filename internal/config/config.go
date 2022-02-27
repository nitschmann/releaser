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
	BranchTitleFormatDefault        string   = "{{if BranchType}}{{ .BranchType }}{{end}} {{ .BranchTitle }}"
	BranchTypesDefault              []string = []string{"bug", "feature", "fix", "hotfix"}
	// Commit config default values
	CommitAllowedWithoutTypeDefault bool     = true
	CommitMessageFormatDefault      string   = "{{if CommitType}}{{ .CommitType | ToTitle }}:{{end}} {{ .CommitMessage }}"
	CommitTypesDefault              []string = []string{"adjustment", "bug", "feature", "fix", "hotfix"}
	// Flag config default values
	FlagRequiredDefault       bool = false
	FlagSkipForBranchDefault  bool = false
	FlagSkipForCommitDefault  bool = false
	FlagSkipForReleaseDefault bool = false
	// Git config default values
	GitExecutableDefault string = "git"
	GitRemoteDefault     string = "origin"
)

// Config has all the relevant settings
type Config struct {
	// Branch specific config fields
	Branch Branch `mapstructure:"branch" yaml:"branch" validate:"required,dive"`
	// Commit specific config fields
	Commit Commit `mapstructure:"commit" yaml:"commit" validate:"required,dive"`
	// Flags specify custom flags for commands
	Flags []Flag `mapstructure:"flags" yaml:"flags" validate:"dive"`
	// Git specific config fields
	Git Git `mapstructure:"git" yaml:"git" validate:"required,dive"`
}

// New returns an new instance of Config with default values
func New() Config {
	return Config{
		Branch: newBranch(),
		Commit: newCommit(),
		Flags:  []Flag{},
		Git:    newGit(),
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
