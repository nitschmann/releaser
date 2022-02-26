package config_test

import (
	"testing"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/stretchr/testify/suite"
)

type gitTestSuite struct {
	suite.Suite
}

func TestGitSuite(t *testing.T) {
	suite.Run(t, new(gitTestSuite))
}

func (s *gitTestSuite) GetExecutable() {
	val := config.Git{}.GetExecutable()
	s.Equal(val, config.GitExecutableDefault)
}

func (s *gitTestSuite) GetRemote() {
	val := config.Git{}.GetRemote()
	s.Equal(val, config.GitRemoteDefault)
}
