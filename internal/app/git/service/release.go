package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/nitschmann/release-log/internal/app/git"
	giturls "github.com/whilp/git-urls"
)

// Service to handle the release
type ReleaseService struct {
	GitRemote  string
	GitRepoUrl string
}

// Returns a new instance of ReleaseService
func NewReleaseService(gitRemote string, gitRepoUrl string) *ReleaseService {
	return &ReleaseService{
		GitRemote:  gitRemote,
		GitRepoUrl: gitRepoUrl,
	}
}

// Returns the title for the new release with the given version tag
func (s ReleaseService) Title(newVersionTag string) string {
	t := time.Now().Format("2006-01-02")

	return fmt.Sprintf("Release %s (%s)", newVersionTag, t)
}

// Creates a compare HTTP URL of the repo for two versions. Returns empty string if latestVersionTag
// is an empty string.
func (s ReleaseService) RepoVersionTagCompareURL(latestVersionTag string, newVersionTag string) (string, error) {
	if latestVersionTag == "" {
		return "", nil
	} else {
		repoUrl, err := s.RepoHttpUrl()
		if err != nil {
			return "", nil
		}

		return strings.Join([]string{repoUrl, "compare", latestVersionTag + "..." + newVersionTag}, "/"), nil
	}
}

// TODO: Maybe move this block into another service or context?
// Returns the git repository URL with http instead of ssh & Co.
func (s ReleaseService) RepoHttpUrl() (string, error) {
	var err error
	var gitRemoteUrl string

	if s.GitRepoUrl != "" {
		gitRemoteUrl = s.GitRepoUrl
	} else {
		gitRemoteUrl, err = git.ExecCommand([]string{"remote", "get-url", s.GitRemote})
		if err != nil {
			return "", err
		}
	}

	url, err := giturls.Parse(gitRemoteUrl)
	if err != nil {
		return "", err
	}

	url.User = nil
	url.Scheme = "http"
	url.Path = strings.Replace(url.Path, ".git", "", 1)

	return url.String(), nil
}
