package config

import "github.com/nitschmann/releaser/internal/data"

// Global is the global config data structure of releaser.
// This is the actual data structure in the YAML configuration file
type Global struct {
	Config
	// Projects which are handled by releaser
	Projects []Project `mapstructure:"projects" yaml:"projects"`
}

// NewGlobal returns an new instance of Global with default values
func NewGlobal() Global {
	return Global{
		Config: New(),
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
