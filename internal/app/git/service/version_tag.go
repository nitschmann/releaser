package service

import (
	"strconv"
	"strings"

	"github.com/nitschmann/release-log/internal/app/git"
)

// Service to handle git (version) tags
type VersionTagService struct {
	defaultFirstVersion string
}

func NewVersionTagService(defaultFirstVersion string) *VersionTagService {
	return &VersionTagService{defaultFirstVersion: defaultFirstVersion}
}

// Builds a new version git tag (returns the defined first version if no git tag is given yet).
// If newVersion par present this one is used
func (s VersionTagService) BuildNew(newVersion string) (string, error) {
	if newVersion == "" {
		versions, err := git.TagList()
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

			return strings.Join(previousVersionParts, "."), nil
		} else {
			return s.defaultFirstVersion, nil
		}
	} else {
		return newVersion, nil
	}
}

// Gets the latest version tag from the list or just return the given latestVersionTag parameter
func (s VersionTagService) LatestVersionTag(latestVersionTag string) (string, error) {
	if latestVersionTag == "" {
		versions, err := git.TagList()
		if err != nil {
			return "", err
		}

		if len(versions) >= 1 {
			return versions[len(versions)-1], nil
		} else {
			return "", nil
		}
	} else {
		return latestVersionTag, nil
	}
}

// Return the previous version tag for a given version tag
func (s VersionTagService) PreviousVersionTag(currentVersionTag string) (string, error) {
	previousVersionParts := strings.Split(strings.TrimPrefix(currentVersionTag, "v"), ".")
	previousVersionPartNum, err := strconv.Atoi(previousVersionParts[len(previousVersionParts)-1])
	if err != nil {
		return "", nil
	}

	previousVersionParts[len(previousVersionParts)-1] = strconv.Itoa(previousVersionPartNum - 1)

	return strings.Join(previousVersionParts, "."), nil
}
