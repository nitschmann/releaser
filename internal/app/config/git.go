package config

import (
	"github.com/go-playground/validator/v10"
)

// Git holds git specific config rules
type Git struct {
	Executable string `mapstructure:"executable" yaml:"executable"`
	Remote     string `mapstructure:"remote" yaml:"remote" validate:"required"`
	RepoURL    string `mapstructure:"repo_url" yaml:"repo_url"`
}

func newGit() *Git {
	return &Git{
		Executable: "git",
		Remote:     "origin",
	}
}

func (g Git) validate() error {
	validate := validator.New()
	err := validate.Struct(g)
	return err
}
