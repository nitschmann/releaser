package git

import (
	"net/url"
	"strings"

	giturls "github.com/whilp/git-urls"
)

// Remote is the interface to manage git remotes
type Remote interface {
	// GetHttpURL returns the URL of the given remote name in HTTP(S) format
	GetHttpURL(name string) (string, error)
	// GetProject returns the project part of the remote remote URL
	GetProject(name string) (string, error)
	// GetProjectByHttpURL extracts the project name of a given repo HTTP url
	GetProjectByHttpURL(httpURL string) (string, error)
	// GetURL returns the URL of the given remote name
	GetURL(name string) (*url.URL, error)
}

type remote struct {
	Git Git
}

// NewRemote returns an instance of Remote interface with default values
func NewRemote(gitObj Git) Remote {
	return remote{Git: gitObj}
}

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

func (r remote) GetProject(name string) (string, error) {
	httpURL, err := r.GetHttpURL(name)
	if err != nil {
		return "", err
	}

	return r.GetProjectByHttpURL(httpURL)
}

func (r remote) GetProjectByHttpURL(httpURL string) (string, error) {
	url, err := url.Parse(httpURL)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(url.Path, "/"), nil
}

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
