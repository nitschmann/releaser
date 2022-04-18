package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"

	"github.com/nitschmann/releaser/pkg/release"
)

// DefaultTokenEnvVar is the default ENV variable name which holds the API token
const DefaultTokenEnvVar string = "GITHUB_API_TOKEN"

// Github upstream for releaser (also supports GitHub enterprise)
type Github struct{}

// New instance of Github
func New() Github {
	return Github{}
}

// APITokenEnvVar specifies which ENV var holds the API token for the GitHub upstream
func (g Github) APITokenEnvVar() string {
	return DefaultTokenEnvVar
}

// Publish the release to GitHub
func (g Github) Publish(
	ctx context.Context,
	apiToken string,
	r *release.Release,
) (*release.UpstreamResult, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	repoURL, err := url.Parse(r.RepoHttpURL)
	if err != nil {
		return nil, err
	}

	var client *github.Client
	if repoURL.Hostname() == "github.com" {
		client = github.NewClient(tc)
	} else {
		repoURL.Path = ""
		baseURL := repoURL.String()

		client, err = github.NewEnterpriseClient(baseURL, baseURL, tc)
		if err != nil {
			return nil, err
		}
	}

	releaseData := &github.RepositoryRelease{
		TagName:         &r.Tag,
		TargetCommitish: &r.Target,
		Name:            &r.Name,
		Body:            &r.Description,
		Draft:           &r.IsDraft,
		Prerelease:      &r.IsPreRelease,
	}

	owner, repo := r.OwnerAndRepo()
	githubRelease, _, err := client.Repositories.CreateRelease(ctx, owner, repo, releaseData)
	if err != nil {
		return nil, err
	}

	return &release.UpstreamResult{ID: githubRelease.ID, URL: githubRelease.URL}, nil
}
