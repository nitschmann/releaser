package service

import (
	"strings"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/nitschmann/release-log/pkg/util/list"
)

// ChangelogService is a service struct to handle git logs
type ChangelogService struct {
	GitService        git.Git
	versionTagService *VersionTagService
}

// NewChangelogService returns a new pointer instance of ChangelogService with the given arguments
func NewChangelogService(gitService git.Git, versionTagService *VersionTagService) *ChangelogService {
	return &ChangelogService{
		GitService:        gitService,
		versionTagService: versionTagService,
	}
}

// ChangelogFromVersionTag returns a list of git logs since the latest given version tag.
// If latestVersionTag parameter is empty all current commits are used.
func (c ChangelogService) ChangelogFromVersionTag(latestVersionTag string) ([]string, error) {
	var logOutput []string
	var gitCmdArgs []string

	if latestVersionTag == "" {
		gitCmdArgs = []string{"log", "--oneline", "--format=format:%s"}
	} else {
		gitCmdArgs = []string{"log", "--oneline", "--format=format:%s", latestVersionTag + "..HEAD"}
	}

	gitCommitLog, err := c.GitService.ExecCommand(gitCmdArgs)
	if err != nil {
		return logOutput, err
	}

	logOutput = strings.Split(gitCommitLog, "\n")

	return list.CleanEmptyStrings(logOutput), nil
}
