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

func (s *globalTestSuite) TestGetConfigWithPresentProjectValues() {
	s.Run("with project param nil and no custom global config", func() {
		globalConfig := config.NewGlobal()
		cfg := globalConfig.GetConfigWithPresentProjectValues(nil)

		s.Equal(cfg.Branch.AllowedWithoutType, &config.BranchAllowedWithoutTypeDefault)
		s.Equal(cfg.Branch.Delimiter, &config.BranchDelimiterDefault)
		s.Equal(cfg.Branch.TitleFormat, &config.BranchTitleFormatDefault)
		s.Equal(cfg.Branch.Types, config.BranchTypesDefault)
		s.Equal(cfg.Commit.AllowedWithoutType, &config.CommitAllowedWithoutTypeDefault)
		s.Equal(cfg.Commit.MessageFormat, &config.CommitMessageFormatDefault)
		s.Equal(cfg.Commit.Types, config.CommitTypesDefault)
		s.Empty(cfg.Flags)
	})

	s.Run("with project param nil and custom global config", func() {
		flags := []config.Flag{{Name: "flag1"}}
		delimiter := "+"
		globalConfig := config.NewGlobal()
		globalConfig.Branch.Delimiter = &delimiter
		globalConfig.Flags = flags

		cfg := globalConfig.GetConfigWithPresentProjectValues(nil)

		s.Equal(cfg.Branch.AllowedWithoutType, &config.BranchAllowedWithoutTypeDefault)
		s.Equal(cfg.Branch.Delimiter, &delimiter)
		s.Equal(cfg.Branch.TitleFormat, &config.BranchTitleFormatDefault)
		s.Equal(cfg.Branch.Types, config.BranchTypesDefault)
		s.Equal(cfg.Commit.AllowedWithoutType, &config.CommitAllowedWithoutTypeDefault)
		s.Equal(cfg.Commit.MessageFormat, &config.CommitMessageFormatDefault)
		s.Equal(cfg.Commit.Types, config.CommitTypesDefault)
		s.NotEmpty(cfg.Flags)
		s.Equal(cfg.Flags, flags)
	})

	s.Run("with project overwrites", func() {
		globalBranchDelimiter := "+"
		globalFlags := []config.Flag{{Name: "flag1"}}
		globalGitRemote := "fork"
		projectBranchDelimiter := "/"
		projectFlags := []config.Flag{{Name: "projectflag1"}, {Name: "projectflag2"}}

		globalConfig := config.NewGlobal()
		globalConfig.Branch.Delimiter = &globalBranchDelimiter
		globalConfig.Flags = globalFlags
		globalConfig.Git.Remote = &globalGitRemote

		project := config.Project{}
		project.Branch.Delimiter = &projectBranchDelimiter
		project.Flags = projectFlags
		globalConfig.Projects = []config.Project{project}

		cfg := globalConfig.GetConfigWithPresentProjectValues(&project)

		s.Equal(cfg.Branch.AllowedWithoutType, &config.BranchAllowedWithoutTypeDefault)
		s.Equal(cfg.Branch.GetDelimiter(), projectBranchDelimiter)
		s.Equal(cfg.Branch.TitleFormat, &config.BranchTitleFormatDefault)
		s.Equal(cfg.Branch.Types, config.BranchTypesDefault)
		s.Equal(cfg.Commit.AllowedWithoutType, &config.CommitAllowedWithoutTypeDefault)
		s.Equal(cfg.Commit.MessageFormat, &config.CommitMessageFormatDefault)
		s.Equal(cfg.Commit.Types, config.CommitTypesDefault)
		s.Equal(cfg.Git.GetRemote(), globalGitRemote)
		s.NotEmpty(cfg.Flags)
		s.Equal(cfg.Flags, project.Flags)
	})
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
