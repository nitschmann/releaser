package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
)

type projectTestSuite struct {
	suite.Suite
}

func TestProjectSuite(t *testing.T) {
	suite.Run(t, new(projectTestSuite))
}

func (s *projectTestSuite) TestMatchesWithPath() {
	project := config.Project{}
	textTemplateValues := data.NewTextTemplateValues()

	s.Run("with global system wildcard", func() {
		project.Paths = []string{"/**/*"}
		match, err := project.MatchesWithPath("/home/user/project", textTemplateValues)

		s.NoError(err)
		s.True(match)
	})

	s.Run("with UserHomeDir template var used", func() {
		textTemplateValues.UserHomeDir = "/home/user"
		project.Paths = []string{"{{ .UserHomeDir }}/**/*"}

		match, err := project.MatchesWithPath("/home/user/code/project", textTemplateValues)

		s.NoError(err)
		s.True(match)
	})

	s.Run("with multiple paths", func() {
		textTemplateValues.UserHomeDir = "/home/user"
		project.Paths = []string{
			"/home/user2/repos/**/*",
			"{{ .UserHomeDir }}/code/**/*",
		}

		match, err := project.MatchesWithPath("/home/user/code/releaser", textTemplateValues)

		s.NoError(err)
		s.True(match)
	})
}
