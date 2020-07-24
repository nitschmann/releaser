package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Application config data struct
type Config struct {
	FirstVersion    string `mapstructure:"first_version" validate="required"`
	GitExecutable   string `mapstructure:"git_executable" validate="required"`
	GitRemote       string `mapstructure:"git_remote" validate="required"`
	GitRepoUrl      string `mapstructure:"git_repo_url"`
	NewVersion      string `mapstructure:"new_version"`
	PreviousVersion string `mapstructure:"previous_version"`
}

var (
	// Global config instance
	Cfg *Config
	// Private vars
	validate *validator.Validate
)

// Returns the current global instance of Config
func Get() *Config {
	return Cfg
}

// Init function to setup config paths and ENV binding. Should just be called once.
func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/release-log")
	viper.AddConfigPath("$HOME/.release-log")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("RELEASE_LOG")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(fmt.Errorf("Error while loading config: %v \n", err))
		}
	}
}

// Loads the config with Viper into a Struct and validates it
func Load() error {
	Cfg = &Config{}
	err := viper.Unmarshal(Cfg)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(Cfg)
	if err != nil {
		return err
	}

	return nil
}

// This sets the default config values. Should just be called once.
func SetDefaultValues() {
	viper.SetDefault("first_version", "v0.0.1")
	viper.SetDefault("git_executable", "git")
	viper.SetDefault("git_remote", "origin")
	viper.SetDefault("git_repo_url", "")
	viper.SetDefault("new_version", "")
	viper.SetDefault("previous_version", "")
}
