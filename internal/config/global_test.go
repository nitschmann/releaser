package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
)

type globalTestSuite struct {
	suite.Suite
}

func TestGlobalSuite(t *testing.T) {
	suite.Run(t, new(globalTestSuite))
}

func (s *globalTestSuite) TestGetProjectConfigByPath() {
	globalConfig := config.NewGlobal()
	textTemplateValues := data.NewTextTemplateValues()

	s.Run("default with empty Projects field", func() {
		globalConfig.Projects = []config.Project{}
		project, err := globalConfig.GetProjectConfigByPath("/home/user/project", textTemplateValues)

		s.NoError(err)
		s.Nil(project)
	})

	s.Run("with match in Projects field entry", func() {
		textTemplateValues.UserHomeDir = "/home/user"
		globalConfig.Projects = []config.Project{
			{
				Paths: []string{"{{ .UserHomeDir }}/repos/**/*"},
			},
		}

		project, err := globalConfig.GetProjectConfigByPath("/home/user/repos/releaser", textTemplateValues)

		s.NoError(err)
		s.Equal(*project, globalConfig.Projects[0])
	})
}
