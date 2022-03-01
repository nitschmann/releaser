package commit_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service/commit"
	"github.com/nitschmann/releaser/test/mock"
)

type createServiceTestSuite struct {
	suite.Suite

	git mock.Git
}

func TestCreateServiceSuite(t *testing.T) {
	suite.Run(t, new(createServiceTestSuite))
}

func (s *createServiceTestSuite) SetupTest() {
	s.git = mock.Git{}
}

func (s *createServiceTestSuite) defaultConfig() config.Config {
	return config.New()
}

func (s *createServiceTestSuite) defaultFlags() map[string]string {
	return make(map[string]string)
}

func (s *createServiceTestSuite) defaultTextTemplateValues() *data.TextTemplateValues {
	return data.NewTextTemplateValues()
}

func (s *createServiceTestSuite) defaultSerivce() commit.CreateService {
	return commit.NewCreateService(s.git)
}

func (s *createServiceTestSuite) TestCall() {
	ctx := context.TODO()

	s.Run("default without type and custom flags", func() {
		commitMessageFormat := "{{if .Type}}{{ .Type }}{{end}} {{ .CommitMessage }}"

		cfg := s.defaultConfig()
		cfg.Commit.MessageFormat = &commitMessageFormat

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()

		commitMessage, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"Add tests",
			"",
			true,
		)

		s.NoError(err)
		s.Equal(commitMessage, "Add tests")
	})

	s.Run("with default type", func() {
		commitMessageFormat := "{{if .Type}}{{ .Type | ToTitle }}:{{end}} {{ .CommitMessage }}"
		commitType := "fix"

		cfg := s.defaultConfig()
		cfg.Commit.MessageFormat = &commitMessageFormat

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()

		commitMessage, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"missing tests  ",
			commitType,
			true,
		)

		s.NoError(err)
		s.Equal(commitMessage, "Fix: missing tests")
	})
}
