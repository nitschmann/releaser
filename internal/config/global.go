package config

import (
	"github.com/spf13/viper"

	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/helper"
)

// Global is the global config data structure of releaser.
// This is the actual data structure in the YAML configuration file
type Global struct {
	Config `mapstructure:",squash"`
	// Projects which are handled by releaser
	Projects []Project `mapstructure:"projects" yaml:"projects"`
}

// NewGlobal returns an new instance of Global with default values
func NewGlobal() *Global {
	return &Global{
		Config: New(),
	}
}

// GetConfigWithPresentProjectValues creates an instance of Config and assigns values of the Project instance to it if present
// Non-nil project param must be part of the Projects field.
func (g Global) GetConfigWithPresentProjectValues(project *Project) Config {
	if project == nil {
		project = &Project{}
	}

	flags := project.Flags
	if len(flags) == 0 {
		flags = g.Flags
	}

	return Config{
		Branch: Branch{
			AllowedWithoutType: helper.BoolPointerOrBackup(project.Branch.AllowedWithoutType, g.Branch.GetAllowedWithoutType()),
			Delimiter:          helper.StringPointerOrBackup(project.Branch.Delimiter, g.Branch.GetDelimiter()),
			NameFormat:        helper.StringPointerOrBackup(project.Branch.NameFormat, g.Branch.GetNameFormat()),
			Types:              helper.StringSliceWithValuesOrBackup(project.Branch.Types, g.Branch.GetTypes()),
		},
		Commit: Commit{
			AllowedWithoutType: helper.BoolPointerOrBackup(project.Commit.AllowedWithoutType, g.Commit.GetAllowedWithoutType()),
			MessageFormat:      helper.StringPointerOrBackup(project.Commit.MessageFormat, g.Commit.GetMessageFormat()),
			Types:              helper.StringSliceWithValuesOrBackup(project.Commit.Types, g.Commit.GetTypes()),
		},
		Git: Git{
			Executable: helper.StringPointerOrBackup(project.Git.Executable, g.Git.GetExecutable()),
			Remote:     helper.StringPointerOrBackup(project.Git.Remote, g.Git.GetRemote()),
		},
		Flags: flags,
	}
}

// GetProjectConfigByPath checks and returns if any of the configured entries in the Projects field matches with the specified path
func (g Global) GetProjectConfigByPath(path string, textTemplateValues *data.TextTemplateValues) (*Project, error) {
	for _, project := range g.Projects {
		match, err := project.MatchesWithPath(path, textTemplateValues)
		if err != nil {
			return nil, err
		}

		if match {
			return &project, nil
		}
	}

	return nil, nil
}

// Load uses viper and unmarshals the YAML config into Global struct
func Load() (*Global, error) {
	globalConfig := &Global{}

	err := viper.Unmarshal(globalConfig)
	if err != nil {
		return nil, err
	}

	return globalConfig, nil
}
