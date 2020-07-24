package git

import (
	"strings"

	"github.com.com/nitschmann/release-log/pkg/util"
)

// List all Git tags in reverse order by vX.X
func TagList() ([]string, error) {
	var tagList []string
	commandOutput, err := ExecCommand([]string{"tag", "-l", "--sort=v:refname"})
	if err != nil {
		return tagList, err
	}

	tagList = strings.Split(commandOutput, "\n")

	return util.CleanList(tagList), nil
}
