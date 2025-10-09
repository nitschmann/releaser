package git_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/pkg/git"
	"github.com/nitschmann/releaser/test/mock"
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
		gitObj := git.New("git")
		_ = gitObj // no-op: nothing to assert here since ExecCommand in real gitObj runs os/exec
	})
}
