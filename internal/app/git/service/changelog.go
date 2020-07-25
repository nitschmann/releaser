package service

import (
	"strings"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/nitschmann/release-log/pkg/util"
)

// Service to handle git logs
type ChangelogService struct {
	versionTagService *VersionTagService
}

func NewChangelogService(versionTagService *VersionTagService) *ChangelogService {
	return &ChangelogService{versionTagService: versionTagService}
}

// Returns a list of git log since the latest given version tag.
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
