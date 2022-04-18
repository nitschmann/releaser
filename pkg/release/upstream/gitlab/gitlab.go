package gitlab

import (
	"context"
	"net/url"
	"path"

	"github.com/xanzy/go-gitlab"

	"github.com/nitschmann/releaser/pkg/release"
)

// DefaultTokenEnvVar is the default ENV variable name which holds the API token
const DefaultTokenEnvVar string = "GITLAB_API_TOKEN"

// Gitlab upstream for releaser
type Gitlab struct{}

// New instance of Gitlab
func New() Gitlab {
	return Gitlab{}
}

// APITokenEnvVar specifies which ENV var holds the API token for the GitHub upstream
func (g Gitlab) APITokenEnvVar() string {
	return DefaultTokenEnvVar
}

// Publish the release to Gitlab
func (g Gitlab) Publish(
	ctx context.Context,
	apiToken string,
	r *release.Release,
) (*release.UpstreamResult, error) {
	repoURL, err := url.Parse(r.RepoHttpURL)
	if err != nil {
		return nil, err
	}

	var gitlabClient *gitlab.Client

	if repoURL.Hostname() == "gitlab.com" {
		gitlabClient, err = gitlab.NewClient(apiToken)
	} else {
		repoURL.Path = ""
		baseURL := repoURL.String()

		gitlabClient, err = gitlab.NewClient(apiToken, gitlab.WithBaseURL(baseURL))
	}

	if err != nil {
		return nil, err
	}

	releaseOptions := &gitlab.CreateReleaseOptions{
		Name:        &r.Name,
		TagName:     &r.Tag,
		Description: &r.Description,
		Ref:         &r.Target,
	}

	owner, repo := r.OwnerAndRepo()
	_, _, err = gitlabClient.Releases.CreateRelease(path.Join(owner, repo), releaseOptions)
	if err != nil {
		return nil, err
	}

	return &release.UpstreamResult{}, nil
}
