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
	Branch Branch `mapstructure:"branch" yaml:"branch" validate:"required,dive"`
	Commit Commit `mapstructure:"commit" yaml:"commit" validate:"required,dive"`
	Flags  []Flag `mapstructure:"flags" yaml:"flags" validate:"dive"`
	Git    Git    `mapstructure:"git" yaml:"git" validate:"required,dive"`
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

// Init loads the config. This just should be called once per runtime
func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./.releaser")
	viper.AddConfigPath("$HOME/.releaser")
	viper.AddConfigPath("/etc/releaser")

	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return err
		}
	}

	return nil
}

// Load uses viper and unmarshals the YAML config into Global
func Load() (*Global, error) {
	globalCfg := NewGlobal()
	cfg := &globalCfg

	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
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
