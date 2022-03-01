package commit

import (
	"context"
	"regexp"
	"strings"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

// CreateService is the service interface to create new git commits
type CreateService interface {
	// Call and execute the process
	Call(
		ctx context.Context,
		cfg config.Config,
		textTemplateValues *data.TextTemplateValues,
		customFlags map[string]string,
		message string,
		commitType string,
		onlyMessage bool,
	) (string, error)
}

type createService struct {
	Git gitPkg.Git
}

// NewCreateService returns an instance which implements the CreateService interface
func NewCreateService(git gitPkg.Git) CreateService {
	return &createService{Git: git}
}

func (s *createService) Call(
	ctx context.Context,
	cfg config.Config,
	textTemplateValues *data.TextTemplateValues,
	customFlags map[string]string,
	message string,
	commitType string,
	onlyMessage bool,
) (string, error) {
	var commitMessage string

	// Check if commitType is valid
	err := service.ValidateType(cfg.Commit.GetTypes(), commitType)
	if err != nil {
		return commitMessage, err
	}

	// Validate custom customFlags
	err = service.ValidateCustomFlags(cfg.GetFlagsForCommit().Names(), customFlags)
	if err != nil {
		return commitMessage, err
	}

	// Assign values for text template
	textTemplateValues.CommitMessage = message
	textTemplateValues.Flags = customFlags
	textTemplateValues.Type = commitType

	// Generate commit message based on template data
	commitMessage, err = textTemplateValues.ParseTemplateString(cfg.Commit.GetMessageFormat())
	if err != nil {
		return commitMessage, err
	}

	// Reduce multiple whitespaces to one
	space := regexp.MustCompile(`\s+`)
	commitMessage = space.ReplaceAllString(commitMessage, " ")
	// Replace all leading and trailing whitespaces
	commitMessage = strings.TrimSpace(commitMessage)

	if !onlyMessage {
		commit := gitPkg.NewCommit(s.Git)
		err = commit.New(commitMessage)
		if err != nil {
			return commitMessage, nil
		}
	}

	return commitMessage, nil
}
