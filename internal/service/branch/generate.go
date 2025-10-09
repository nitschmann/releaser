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

	// Sanitize invalid git branch name characters/sequences by replacing with delimiter
	branchName = sanitizeBranchName(branchName, delimiter)

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

// sanitizeBranchName replaces characters and sequences that are invalid for git branch names
// with the provided delimiter and performs final cleanups to produce a valid candidate.
// See `git check-ref-format --branch` rules.
func sanitizeBranchName(name string, delimiter string) string {
	if name == "" {
		return name
	}

	// Replace ASCII control characters (including DEL) with delimiter
	ctrl := regexp.MustCompile(`[\x00-\x1F\x7F]`)
	name = ctrl.ReplaceAllString(name, delimiter)

	// Replace sequences that are explicitly forbidden
	name = strings.ReplaceAll(name, "..", delimiter)
	name = strings.ReplaceAll(name, "@{", delimiter)

	// Replace characters not allowed in ref names: space, ~, ^, :, ?, *, [, \
	// Note: slash '/' is allowed inside branch names, we keep it
	invalidChars := regexp.MustCompile(`[\s~^:?*\[\\\]\{\}]`)
	name = invalidChars.ReplaceAllString(name, delimiter)

	// Collapse multiple slashes
	multiSlash := regexp.MustCompile(`/+`)
	name = multiSlash.ReplaceAllString(name, "/")

	// Disallow path components '.' or '..' by replacing them with delimiter
	dotComponent := regexp.MustCompile(`(^|/)\.{1,2}($|/)`)
	name = dotComponent.ReplaceAllString(name, "$1"+delimiter+"$2")

	// Cannot end with a dot
	name = strings.TrimRight(name, ".")

	// Cannot end with ".lock"
	name = strings.TrimSuffix(name, ".lock")

	// Trim leading/trailing slashes
	name = strings.Trim(name, "/")

	// Collapse multiple delimiters
	if delimiter != "" {
		// Escape delimiter for regex class
		esc := regexp.QuoteMeta(delimiter)
		multiDelim := regexp.MustCompile(esc + `+`)
		name = multiDelim.ReplaceAllString(name, delimiter)

		// Ensure name does not start or end with delimiter
		name = strings.TrimPrefix(name, delimiter)
		name = strings.TrimSuffix(name, delimiter)

		// In case trimming the delimiter exposes a leading or trailing slash again
		name = strings.Trim(name, "/")
	}

	// As a last resort, if empty, return a safe fallback
	if name == "" {
		name = "branch"
	}

	return name
}
