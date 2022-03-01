package branch_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service/branch"
	"github.com/nitschmann/releaser/test/mock"
)

type generateServiceTestSuite struct {
	suite.Suite

	git mock.Git
}

func TestGenerateServiceSuite(t *testing.T) {
	suite.Run(t, new(generateServiceTestSuite))
}

func (s *generateServiceTestSuite) SetupTest() {
	s.git = mock.Git{}
}

func (s *generateServiceTestSuite) defaultConfig() config.Config {
	return config.New()
}

func (s *generateServiceTestSuite) defaultFlags() map[string]string {
	return make(map[string]string)
}

func (s *generateServiceTestSuite) defaultTextTemplateValues() *data.TextTemplateValues {
	return data.NewTextTemplateValues()
}

func (s *generateServiceTestSuite) defaultSerivce() branch.GenerateService {
	return branch.NewGenerateService(s.git)
}

func (s *generateServiceTestSuite) TestCall() {
	ctx := context.TODO()

	s.Run("default without type and custom flags", func() {
		delimiter := "-"
		branchNameFormat := "{{ .BranchName }}"

		cfg := s.defaultConfig()
		cfg.Branch.Delimiter = &delimiter
		cfg.Branch.NameFormat = &branchNameFormat

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()

		branchName, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"Feature 1",
			"",
			false,
		)

		s.NoError(err)
		s.Equal(branchName, "feature-1")
	})

	s.Run("with valid default branch type given", func() {
		delimiter := "-"
		branchNameFormat := "{{if .Type}}{{ .Type }}-{{end}}{{ .BranchName }}"
		branchType := "fix"

		cfg := s.defaultConfig()
		cfg.Branch.Delimiter = &delimiter
		cfg.Branch.NameFormat = &branchNameFormat

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()

		branchName, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"Endpoint",
			branchType,
			false,
		)

		s.NoError(err)
		s.Equal(branchName, "fix-endpoint")
	})

	s.Run("with valid custom branch type given", func() {
		delimiter := "-"
		branchNameFormat := "{{if .Type}}{{ .Type }}-{{end}}{{ .BranchName }}"
		customBranchType := "custombranch"

		cfg := s.defaultConfig()
		cfg.Branch.Delimiter = &delimiter
		cfg.Branch.NameFormat = &branchNameFormat
		cfg.Branch.Types = append(cfg.Branch.Types, customBranchType)

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()

		branchName, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"New Endpoint",
			customBranchType,
			false,
		)

		s.NoError(err)
		s.Equal(branchName, fmt.Sprintf("%s-%s", customBranchType, "new-endpoint"))
	})

	s.Run("with custom flag set", func() {
		delimiter := "-"
		branchNameFormat := "{{if .Flags.issue}}issue {{ .Flags.issue }}{{end}} {{if .Type}}{{ .Type }}{{end}} {{ .BranchName }}"
		customBranchType := "fix"
		customIssue := "1"

		cfg := s.defaultConfig()
		cfg.Flags = []config.Flag{{Name: "issue"}}
		cfg.Branch.Delimiter = &delimiter
		cfg.Branch.NameFormat = &branchNameFormat
		cfg.Branch.Types = append(cfg.Branch.Types, customBranchType)

		textTemplateValues := s.defaultTextTemplateValues()
		flags := s.defaultFlags()
		flags["issue"] = customIssue

		branchName, err := s.defaultSerivce().Call(
			ctx,
			cfg,
			textTemplateValues,
			flags,
			"Endpoint",
			customBranchType,
			false,
		)

		s.NoError(err)
		s.Equal(branchName, fmt.Sprintf("issue-%s-%s-%s", customIssue, customBranchType, "endpoint"))
	})
}
