package service

import (
	"strconv"
	"strings"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/nitschmann/release-log/internal/app/git/tag"
)

// VersionTagService is a service struct to handle version tags
type VersionTagService struct {
	defaultFirstVersion string
	GitService          git.Git
}

// NewVersionTagService returns a new pointer instance of VersionTagService with the given arguments
func NewVersionTagService(g git.Git, defaultFirstVersion string) *VersionTagService {
	return &VersionTagService{
		defaultFirstVersion: defaultFirstVersion,
		GitService:          g,
	}
}

// CreateNew builds a new version git tag or returns the config defined first version if no git tag is given yet.
// If newVersion parameter is present this one is used instead.
func (s VersionTagService) CreateNew(newVersion string) (string, error) {
	if newVersion == "" {
		gitTag := tag.New(s.GitService)
		versions, err := gitTag.SortedList("v:refname")
		if err != nil {
			return "", err
		}

		if len(versions) >= 1 {
			previousVersion := versions[len(versions)-1]
			previousVersionParts := strings.Split(strings.TrimPrefix(previousVersion, "v"), ".")
			previousVersionPartNum, err := strconv.Atoi(previousVersionParts[len(previousVersionParts)-1])
			if err != nil {
				return "", nil
			}

			newVersionNum := previousVersionPartNum + 1
			previousVersionParts[len(previousVersionParts)-1] = strconv.Itoa(newVersionNum)

			return "v" + strings.Join(previousVersionParts, "."), nil
		}

		return s.defaultFirstVersion, nil
	}

	return newVersion, nil
}

// LatestVersionTag gets and returns the latest version tag from the list through git or
// just returns the latestVersionTag parameter (if given)
func (s VersionTagService) LatestVersionTag(latestVersionTag string) (string, error) {
	if latestVersionTag == "" {
		gitTag := tag.New(s.GitService)
		versions, err := gitTag.SortedList("v:refname")
		if err != nil {
			return "", err
		}

		if len(versions) >= 1 {
			return versions[len(versions)-1], nil
		}

		return "", nil
	}

	return latestVersionTag, nil
}
