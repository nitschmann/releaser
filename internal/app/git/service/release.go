package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/nitschmann/releaser/internal/app/git"
	giturls "github.com/whilp/git-urls"
)

// ReleaseService is a service struct to handle release (logs)
type ReleaseService struct {
	GitRemote  string
	GitRepoURL string
	GitService git.Git
}

// NewReleaseService returns a new pointer instance of ReleaseService with the given arguments
func NewReleaseService(gitService git.Git, gitRemote string, gitRepoURL string) *ReleaseService {
	return &ReleaseService{
		GitService: gitService,
		GitRemote:  gitRemote,
		GitRepoURL: gitRepoURL,
	}
}

// Title returns the title for the new release with the given version tag
func (s ReleaseService) Title(newVersionTag string) string {
	t := time.Now().Format("2006-01-02")

	return fmt.Sprintf("Release %s (%s)", newVersionTag, t)
}

// RepoVersionTagCompareURL creates the release compare HTTP URL for two versions.
// It returns an empty string if the latestVersionTag parameter is not present.
func (s ReleaseService) RepoVersionTagCompareURL(latestVersionTag string, newVersionTag string) (string, error) {
	if latestVersionTag == "" {
		return "", nil
	}

	repoURL, err := s.RepoHTTPURL()
	if err != nil {
		return "", nil
	}

	return strings.Join([]string{repoURL, "compare", latestVersionTag + "..." + newVersionTag}, "/"), nil
}

// RepoHTTPURL returns the git repository URL with http instead of ssh & Co.
// TODO: Maybe move this block into another service or context?
func (s ReleaseService) RepoHTTPURL() (string, error) {
	var err error
	var gitRemoteURL string

	if s.GitRepoURL != "" {
		gitRemoteURL = s.GitRepoURL
	} else {
		gitRemoteURL, err = s.GitService.ExecCommand([]string{"remote", "get-url", s.GitRemote})
		if err != nil {
			return "", err
		}
	}

	url, err := giturls.Parse(gitRemoteURL)
	if err != nil {
		return "", err
	}

	url.User = nil
	url.Scheme = "http"
	url.Path = strings.Replace(url.Path, ".git", "", 1)

	return url.String(), nil
}
