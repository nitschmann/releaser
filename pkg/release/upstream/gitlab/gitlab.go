package gitlab

import (
	"context"
	"net/url"
	"strings"

	"github.com/nitschmann/releaser/pkg/release"
)

// DefaultTokenEnvVar is the default ENV variable name which holds the API token
const DefaultTokenEnvVar string = "GITLAB_API_TOKEN"

// Gitlab upstream for releaser (also supports GitHub enterprise)
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

	ownerAndRepo := strings.Split(r.RepoName, "/")
	owner := ownerAndRepo[0]
	repo := ownerAndRepo[1]

	return nil, nil
}
