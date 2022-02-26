package config

import (
	"github.com/mattn/go-zglob"

	"github.com/nitschmann/releaser/internal/data"
)

// Project is the project (path) specific config data structure
type Project struct {
	Config
	// Paths of the system to which these config settings should apply to. Supports wildcards
	Paths []string `mapstructure:"paths" yaml:"paths"`
}

func newProject() Project {
	return Project{
		Config: New(),
	}
}

// MatchesWithPath checks if a given path string matches with the configured Path field
func (p Project) MatchesWithPath(pp string, textTemplateValues *data.TextTemplateValues) (bool, error) {
	var match bool

	parsedPaths, err := p.parsedPaths(textTemplateValues)
	if err != nil {
		return match, err
	}

	for _, parsedPath := range parsedPaths {
		match, err = zglob.Match(parsedPath, pp)
		if err != nil {
			return match, err
		}

		if !match {
			continue
		}

		return match, nil
	}

	return match, nil
}

func (p Project) parsedPaths(textTemplateValues *data.TextTemplateValues) ([]string, error) {
	var paths []string

	for _, pp := range p.Paths {
		parsedPath, err := textTemplateValues.ParseTemplateString(pp)
		if err != nil {
			return paths, err
		}

		paths = append(paths, parsedPath)
	}

	return paths, nil
}
