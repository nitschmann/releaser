package git

import (
	"net/url"
	"strings"

	giturls "github.com/whilp/git-urls"
)

// Remote is the interface to manage git remotes
type Remote interface {
	GetHttpURL(name string) (string, error)
	GetURL(name string) (*url.URL, error)
}

type remote struct {
	Git Git
}

// NewRemote returns an instance of Remote interface with default values
func NewRemote(gitObj Git) Remote {
	return remote{Git: gitObj}
}

// GetHttpURL returns the URL of the given remote name in HTTP(S) format
func (r remote) GetHttpURL(name string) (string, error) {
	url, err := r.GetURL(name)
	if err != nil {
		return "", err
	}

	url.User = nil
	url.Path = strings.Replace(url.Path, ".git", "", 1)
	if url.Scheme != "http" {
		url.Scheme = "https"
	}

	return url.String(), nil
}

// GetURL returns the URL of the given remote name
func (r remote) GetURL(name string) (*url.URL, error) {
	var gitURL *url.URL

	gitRemoteURL, _, err := r.Git.ExecCommand([]string{"ls-remote", "--get-url", name})
	if err != nil {
		return gitURL, nil
	}

	gitURL, err = giturls.Parse(gitRemoteURL)
	if err != nil {
		return gitURL, nil
	}

	return gitURL, nil
}
