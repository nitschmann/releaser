package upstream

import (
	"context"

	"github.com/nitschmann/releaser/pkg/release"
)

// Upstream is the interface to pushlish releases
type Upstream interface {
	// APITokenEnvVar specifies which ENV var holds the API token for the upstream
	APITokenEnvVar() string
	// Publish  a release to the upstream
	Publish(ctx context.Context, apiToken string, r *release.Release) (*release.UpstreamResult, error)
}
