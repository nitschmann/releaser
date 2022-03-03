package git_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/pkg/git"
	"github.com/nitschmann/releaser/test/mock"
)

type remoteTestSuite struct {
	suite.Suite

	git mock.Git
}

func TestRemoteSuite(t *testing.T) {
	suite.Run(t, new(remoteTestSuite))
}

func (s *remoteTestSuite) SetupTest() {
	s.git = mock.Git{}
}

func (s *remoteTestSuite) TearDownTest() {
	s.git.AssertExpectations(s.T())
}

func (s *remoteTestSuite) TestGetHttpURL() {
	remoteName := "origin"
	gitCmdArgs := []string{"ls-remote", "--get-url", remoteName}

	s.Run("with SSH git remote URL", func() {
		expectedHost := "github.com"
		expectedPath := "nitschmann/releaser"

		returnedRemoteURLStr := fmt.Sprintf("git@%s:%s.git", expectedHost, expectedPath)
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		url, err := remote.GetHttpURL(remoteName)

		s.NoError(err)
		s.Equal(url, fmt.Sprintf("https://%s/%s", expectedHost, expectedPath))
	})

	s.Run("with HTTPS git remote URL", func() {
		expectedHost := "github.com"
		expectedPath := "nitschmann/releaser"

		returnedRemoteURLStr := fmt.Sprintf("https://%s/%s.git", expectedHost, expectedPath)
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		url, err := remote.GetHttpURL(remoteName)

		s.NoError(err)
		s.Equal(url, fmt.Sprintf("https://%s/%s", expectedHost, expectedPath))
	})

	s.Run("with HTTP git remote URL", func() {
		expectedHost := "github.com"
		expectedPath := "nitschmann/releaser"

		returnedRemoteURLStr := fmt.Sprintf("http://%s/%s.git", expectedHost, expectedPath)
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		url, err := remote.GetHttpURL(remoteName)

		s.NoError(err)
		s.Equal(url, fmt.Sprintf("http://%s/%s", expectedHost, expectedPath))
	})

}

func (s *remoteTestSuite) TestGetProject() {
	remoteName := "origin"
	gitCmdArgs := []string{"ls-remote", "--get-url", remoteName}
	returnedRemoteURLStr := "git@github.com:nitschmann/releaser.git"
	s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

	remote := git.NewRemote(s.git)
	project, err := remote.GetProject(remoteName)

	s.NoError(err)
	s.Equal(project, "nitschmann/releaser")
}

func (s *remoteTestSuite) TestGetURL() {
	remoteName := "origin"
	gitCmdArgs := []string{"ls-remote", "--get-url", remoteName}

	s.Run("with SSH github.com host", func() {
		returnedRemoteURLStr := "git@github.com:nitschmann/releaser.git"
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		remoteURL, err := remote.GetURL(remoteName)

		s.NoError(err)
		s.NotNil(remoteURL)
		s.NotNil(remoteURL.User)
		s.Equal(remoteURL.Host, "github.com")
		s.Equal(remoteURL.Scheme, "ssh")
	})

	s.Run("with HTTPS github.com host", func() {
		returnedRemoteURLStr := "https://github.com/nitschmann/releaser.git"
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		remoteURL, err := remote.GetURL(remoteName)

		s.NoError(err)
		s.NotNil(remoteURL)
		s.Nil(remoteURL.User)
		s.Equal(remoteURL.Host, "github.com")
		s.Equal(remoteURL.Scheme, "https")
	})

	s.Run("with SSH GitHub enterprise host", func() {
		host := "github.company.org"
		returnedRemoteURLStr := fmt.Sprintf("git@%s:project/repo.git", host)
		s.git.On("ExecCommand", gitCmdArgs).Once().Return(returnedRemoteURLStr, 0, nil)

		remote := git.NewRemote(s.git)
		remoteURL, err := remote.GetURL(remoteName)

		s.NoError(err)
		s.NotNil(remoteURL)
		s.NotNil(remoteURL.User)
		s.Equal(remoteURL.Host, host)
		s.Equal(remoteURL.Scheme, "ssh")
	})
}
