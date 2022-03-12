package release

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
	"github.com/nitschmann/releaser/pkg/release"
)

// GenerateService is the service interface to generate releases
type GenerateService interface {
	// Call and execute the process
	Call(
		ctx context.Context,
		cfg config.Config,
		textTemplateValues *data.TextTemplateValues,
		customFlags map[string]string,
		firstTag string,
		gitRemote string,
		tag string,
		target string,
	) (*release.Release, error)
}

type generateService struct {
	Git gitPkg.Git

	gitCommit gitPkg.Commit
	gitRemote gitPkg.Remote
	gitTag    gitPkg.Tag
}

func NewGenerateService(git gitPkg.Git) GenerateService {
	return &generateService{
		Git:       git,
		gitCommit: gitPkg.NewCommit(git),
		gitRemote: gitPkg.NewRemote(git),
		gitTag:    gitPkg.NewTag(git),
	}
}

func (s *generateService) Call(
	ctx context.Context,
	cfg config.Config,
	textTemplateValues *data.TextTemplateValues,
	customFlags map[string]string,
	firstTag string,
	gitRemote string,
	tag string,
	target string,
) (*release.Release, error) {
	var result *release.Release = &release.Release{}

	// Validate custom customFlags
	err := service.ValidateCustomFlags(cfg.GetFlagsForRelease().Names(), customFlags)
	if err != nil {
		return nil, err
	}

	textTemplateValues.Flags = customFlags

	// Set values for template and release object
	err = s.setValues(
		firstTag,
		gitRemote,
		tag,
		target,
		textTemplateValues,
		result,
	)
	if err != nil {
		return nil, err
	}

	// Generate release name
	releaseName, err := textTemplateValues.ParseTemplateString(cfg.Release.GetNameFormat())
	if err != nil {
		return nil, err
	}

	space := regexp.MustCompile(`\s+`)
	releaseName = space.ReplaceAllString(releaseName, " ")
	// Replace all leading and trailing whitespaces
	releaseName = strings.TrimSpace(releaseName)

	// Generate release description log
	releaseDescription, err := textTemplateValues.ParseTemplateString(cfg.Release.GetDescriptionFormat())
	if err != nil {
		return nil, err
	}
	releaseDescription = strings.TrimSpace(releaseDescription)

	result.Description = releaseDescription
	result.Name = releaseName

	return result, nil
}

func (s *generateService) setValues(
	firstTag string,
	gitRemote string,
	tag string,
	target string,
	textTemplateValues *data.TextTemplateValues,
	r *release.Release,
) error {
	gitTagList, err := s.gitTag.List()
	if err != nil {
		return err
	}
	gitTagLatest := gitTagList.Latest()

	if tag == "" {
		tag, err = gitTagList.GenerateNew()
		if err != nil {
			return err
		}

		if tag == "" {
			tag = firstTag
		}
	} else if gitTagList.Includes(tag) {
		return fmt.Errorf("tag '%s' is already present", tag)
	}

	gitCommitLogs, err := s.gitCommit.LogsBetweenVersions(gitTagLatest, gitRemote+"/"+target)
	if err != nil {
		return err
	}

	gitRepoHttpURL, err := s.gitRemote.GetHttpURL(gitRemote)
	if err != nil {
		return err
	}

	gitRepoProjectName, err := s.gitRemote.GetProjectByHttpURL(gitRepoHttpURL)
	if err != nil {
		return err
	}

	textTemplateValues.GitRepoHttpURL = gitRepoHttpURL
	textTemplateValues.GitRemote = gitRemote
	textTemplateValues.GitCommitLogs = gitCommitLogs
	textTemplateValues.ReleaseTag = tag
	textTemplateValues.ReleaseTarget = target

	r.Tag = tag
	r.Target = target
	r.RepoHttpURL = gitRepoHttpURL
	r.RepoName = gitRepoProjectName

	return nil
}
