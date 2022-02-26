package git_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	mock "github.com/nitschmann/releaser/test/mock"
)

type gitTestSuite struct {
	suite.Suite

	git mock.Git
}

func TestGitSuite(t *testing.T) {
	suite.Run(t, new(gitTestSuite))
}

func (s *gitTestSuite) SetupTest() {
	s.git = mock.Git{}
}

func (s *gitTestSuite) TearDownTest() {
	s.git.AssertExpectations(s.T())
}

func (s *gitTestSuite) TestExecCommand() {
	s.Run("happy path", func() {
		args := []string{"log"}
		expectedOutput := "fatal: your current branch 'main' does not have any commits yet"
		expectedCode := 0
		s.git.On("ExecCommand", args).Once().Return(expectedOutput, expectedCode, nil)

		output, code, err := s.git.ExecCommand(args)

		s.NoError(err)
		s.Equal(code, expectedCode)
		s.Equal(output, expectedOutput)
	})
}
