package service

import (
	"strings"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/nitschmann/release-log/pkg/util"
)

// ChangelogService is a service struct to handle git logs
type ChangelogService struct {
	versionTagService *VersionTagService
}

// NewChangelogService returns a new pointer instance of ChangelogService with the given arguments
func NewChangelogService(versionTagService *VersionTagService) *ChangelogService {
	return &ChangelogService{versionTagService: versionTagService}
}

// ChangelogFromVersionTag returns a list of git logs since the latest given version tag.
// If latestVersionTag parameter is empty all current commits are used.
func (s ChangelogService) ChangelogFromVersionTag(latestVersionTag string) ([]string, error) {
	var logOutput []string
	var gitCmdArgs []string

	if latestVersionTag == "" {
		gitCmdArgs = []string{"log", "--oneline", "--format=format:%s"}
	} else {
		gitCmdArgs = []string{"log", "--oneline", "--format=format:%s", latestVersionTag + "..HEAD"}
	}

	gitCommitLog, err := git.ExecCommand(gitCmdArgs)
	if err != nil {
		return logOutput, err
	}

	logOutput = strings.Split(gitCommitLog, "\n")

	return util.CleanList(logOutput), nil
}
