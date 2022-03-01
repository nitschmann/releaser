package branch

import (
	"context"
	"regexp"
	"strings"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

// GenerateService is the service interface to generate and checkout new git branches with name
type GenerateService interface {
	// Call and execute the process
	Call(
		ctx context.Context,
		cfg config.Config,
		textTemplateValues *data.TextTemplateValues,
		customFlags map[string]string,
		name string,
		branchType string,
		checkout bool,
	) (string, error)
}

type generateService struct {
	Git gitPkg.Git
}

// NewGenerateService returns an instance which implements the GenerateService interface
func NewGenerateService(git gitPkg.Git) GenerateService {
	return &generateService{Git: git}
}

func (s *generateService) Call(
	ctx context.Context,
	cfg config.Config,
	textTemplateValues *data.TextTemplateValues,
	customFlags map[string]string,
	name string,
	branchType string,
	checkout bool,
) (string, error) {
	var branchName string

	// Check if branchType is valid if given
	err := service.ValidateType(cfg.Branch.GetTypes(), branchType, true)
	if err != nil {
		return branchName, err
	}

	// Validate custom customFlags
	err = service.ValidateCustomFlags(cfg.GetFlagsForBranch().Names(), customFlags)
	if err != nil {
		return branchName, err
	}

	// Assign values for text template
	textTemplateValues.BranchName = strings.ToLower(name)
	textTemplateValues.Flags = customFlags
	textTemplateValues.Type = branchType

	// Generate basic branch name based on template data
	branchName, err = textTemplateValues.ParseTemplateString(cfg.Branch.GetNameFormat())
	if err != nil {
		return branchName, err
	}

	// Reduce multiple whitespaces to one
	space := regexp.MustCompile(`\s+`)
	branchName = space.ReplaceAllString(branchName, " ")
	// Replace rest of spaces with delimiter
	delimiter := cfg.Branch.GetDelimiter()
	branchName = space.ReplaceAllString(branchName, delimiter)
	// Replace all leading and trailing whitespaces
	branchName = strings.TrimSpace(branchName)
	// Ensure string does not start and end with delimiter
	branchName = strings.TrimPrefix(branchName, delimiter)
	branchName = strings.TrimSuffix(branchName, delimiter)

	if checkout {
		err = s.checkoutBranch(branchName)
		if err != nil {
			return "", err
		}
	}

	return branchName, nil
}

func (s *generateService) checkoutBranch(branchName string) error {
	branch := gitPkg.NewBranch(s.Git)

	err := branch.ValidateName(branchName)
	if err != nil {
		return err
	}

	return branch.Checkout(branchName)
}
