package release

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/helper"
	"github.com/nitschmann/releaser/internal/service"
	"github.com/nitschmann/releaser/pkg/release"
	upstreamPkg "github.com/nitschmann/releaser/pkg/release/upstream"
)

// PublishService is the service interface to publish a release to a specified upstream
type PublishService interface {
	// Call and execute the process
	Call(
		ctx context.Context,
		cfg config.Config,
		releaseObj *release.Release,
		upstreamName string,
		apiToken string,
		autoYes bool,
	) (*release.UpstreamResult, error)
}

type publishService struct{}

func NewPublishService() PublishService {
	return &publishService{}
}

func (s *publishService) Call(
	ctx context.Context,
	cfg config.Config,
	releaseObj *release.Release,
	upstreamName string,
	apiToken string,
	autoYes bool,
) (*release.UpstreamResult, error) {
	var err error

	// Validate upstream name
	upstreamNames := upstreamPkg.GetRegistry().Names()
	if !helper.StringSliceIncludesElement(upstreamNames, upstreamName) {
		return nil, apperror.NewInvalidUpstreamNameError(upstreamName, upstreamNames)
	}

	upstream := upstreamPkg.GetRegistry().Get(upstreamName)

	if !autoYes {
		err = s.printReleaseConfirm(releaseObj)
		if err != nil {
			return nil, err
		}
	}

	if apiToken == "" {
		apiToken, err = s.getAPITokenByEnv(cfg, upstreamName, upstream)
		if err != nil {
			return nil, err
		}
	}

	return upstream.Publish(ctx, apiToken, releaseObj)
}

func (s *publishService) getAPITokenByEnv(
	cfg config.Config,
	upstreamName string,
	upstream upstreamPkg.Upstream,
) (string, error) {
	var (
		apiToken   string
		envVarName string
	)

	err := godotenv.Load()
	if err != nil {
		return apiToken, nil
	}

	upstreamConfig, ok := cfg.Release.GetUpstreams()[upstreamName]
	if ok {
		envVarName = *upstreamConfig.APITokenEnvVar
	} else {
		envVarName = upstream.APITokenEnvVar()
	}

	if envVarName == "" {
		return apiToken, apperror.NewMissingUpstreamAPITokenEnvVarError(upstreamName)
	}

	apiToken = os.Getenv(envVarName)

	return apiToken, nil
}

func (s *publishService) printReleaseConfirm(releaseObj *release.Release) error {
	fmt.Println("--- Release Summary:")
	fmt.Println("\nRepository URL: " + releaseObj.RepoHttpURL)
	fmt.Println("Target: " + releaseObj.Target)
	fmt.Println("Tag: " + releaseObj.Tag)
	fmt.Println("Name: " + releaseObj.Name)
	fmt.Printf("Draft: %t\n", releaseObj.IsDraft)
	fmt.Printf("Pre-Release: %t\n", releaseObj.IsPreRelease)
	fmt.Println("Description:")
	fmt.Println("\n" + releaseObj.Description + "\n")

	return service.PromptYesOrNoWithExpectedYes("Do you want to publish this release?")
}
