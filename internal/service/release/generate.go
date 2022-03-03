package release

import (
	"context"
	"fmt"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/data"
	"github.com/nitschmann/releaser/internal/service"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
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
	) error
}

type generateService struct {
	Git gitPkg.Git

	gitCommit gitPkg.Commit
	gitTag    gitPkg.Tag
}

func NewGenerateService(git gitPkg.Git) GenerateService {
	return &generateService{
		Git:       git,
		gitCommit: gitPkg.NewCommit(git),
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
) error {
	// Validate custom customFlags
	err := service.ValidateCustomFlags(cfg.GetFlagsForRelease().Names(), customFlags)
	if err != nil {
		return err
	}

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

	textTemplateValues.GitRemote = gitRemote
	textTemplateValues.GitCommitLogs = gitCommitLogs
	textTemplateValues.Flags = customFlags
	textTemplateValues.ReleaseTag = tag
	textTemplateValues.ReleaseTarget = target

	releaseName, err := textTemplateValues.ParseTemplateString(cfg.Release.GetNameFormat())
	if err != nil {
		return err
	}

	return nil
}
