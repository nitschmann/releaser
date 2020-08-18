package tag

import (
	"strings"

	"github.com/nitschmann/release-log/pkg/util/list"
)

// List returns a list of git tags
func (t Obj) List() ([]string, error) {
	var tagList []string
	tagList, err := t.ListWithArgs([]string{})
	if err != nil {
		return tagList, err
	}

	return tagList, nil
}

// ListWithArgs returns a list of git tags and allows to pass optional args to the git command
func (t Obj) ListWithArgs(args []string) ([]string, error) {
	var tagList []string
	commandOutput, err := t.Git.ExecCommand(append([]string{"tag", "-l"}, args...))
	if err != nil {
		return tagList, err
	}

	tagList = strings.Split(commandOutput, "\n")

	return list.CleanEmptyStrings(tagList), nil
}

// SortedList returns list of git tags and passes the --sort= flag to the git command. This needs to be set by parameter.
func (t Obj) SortedList(sortKey string) ([]string, error) {
	var tagList []string
	tagList, err := t.ListWithArgs([]string{"--sort=" + sortKey})
	if err != nil {
		return tagList, err
	}

	return tagList, nil
}
