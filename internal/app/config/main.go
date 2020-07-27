package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is the application configuration data struct
type Config struct {
	FirstVersion  string `mapstructure:"first_version" validate="required"`
	GitExecutable string `mapstructure:"git_executable" validate="required"`
	GitRemote     string `mapstructure:"git_remote" validate="required"`
	GitRepoURL    string `mapstructure:"git_repo_url"`
	LatestVersion string `mapstructure:"latest_version"`
	NewVersion    string `mapstructure:"new_version"`
}

var (
	// Cfg is the global instance of Config
	Cfg *Config
	// Private vars
	validate *validator.Validate
)

// Get returns Cfg var
func Get() *Config {
	return Cfg
}

// Init function is to setup config paths and their ENV binding. Should just be called once.
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
			panic(fmt.Errorf("error while loading config: %v", err))
		}
	}
}

// Load uses viper, loads it and sets the output in Cfg
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

// SetDefaultValues sets the default values for the config keys managed with viper.
func SetDefaultValues() {
	viper.SetDefault("first_version", "v0.0.1")
	viper.SetDefault("git_executable", "git")
	viper.SetDefault("git_remote", "origin")
	viper.SetDefault("git_repo_url", "")
	viper.SetDefault("new_version", "")
	viper.SetDefault("previous_version", "")
}
