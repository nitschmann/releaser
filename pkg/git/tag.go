package git

import (
	"strconv"
	"strings"
)

type Tag interface {
	// GenerateNew a new tag or use the backupTag if no tag exists yet
	GenerateNew(backupTag string) (string, error)
	// Latest tag
	Latest() (string, error)
	// List all available git tags accross all branches sorted by date commited
	List() (TagList, error)
}

type tag struct {
	Git Git
}

// NewTag returns an new instance of Tag interface with default values
func NewTag(gitObj Git) Tag {
	return tag{Git: gitObj}
}

func (t tag) GenerateNew(backupTag string) (string, error) {
	tagList, err := t.List()
	if err != nil {
		return "", err
	}

	newTag, err := tagList.GenerateNew()
	if err != nil {
		return "", err
	}

	if newTag == "" {
		return backupTag, nil
	}

	return newTag, nil
}

func (t tag) Latest() (string, error) {
	list, err := t.List()
	if err != nil {
		return "", err
	}

	return list.Latest(), nil
}

func (t tag) List() (TagList, error) {
	var list TagList

	_, _, err := t.Git.ExecCommand([]string{"fetch", "--all", "--quiet"})
	if err != nil {
		return list, err
	}

	gitTagList, _, err := t.Git.ExecCommand([]string{"tag", "--list", "--sort=committerdate"})
	if err != nil {
		return list, err
	}

	return cleanEmptyEntriesFromStringSlice(strings.Split(gitTagList, "\n")), nil
}

// TagList is a list of git tags
type TagList []string

func (tl TagList) GenerateNew() (string, error) {
	latestTag := tl.Latest()
	if latestTag == "" {
		return "", nil
	}

	var newTag string

	latestTagParts := strings.Split(latestTag, ".")
	latestTagPartVersionNum, err := strconv.Atoi(latestTagParts[len(latestTagParts)-1])
	if err != nil {
		return newTag, err
	}

	newTagParts := latestTagParts
	newTagParts[len(newTagParts)-1] = strconv.Itoa(latestTagPartVersionNum + 1)

	newTag = strings.Join(newTagParts, ".")

	return newTag, nil
}

// Includes the specified element
func (tl TagList) Includes(element string) bool {
	for _, tag := range tl {
		if tag == element {
			return true
		}
	}

	return false
}

// Latest entry of the list
func (tl TagList) Latest() string {
	if len(tl) == 0 {
		return ""
	}

	return tl[len(tl)-1]
}
