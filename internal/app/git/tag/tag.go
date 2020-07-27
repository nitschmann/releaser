package tag

import "github.com/nitschmann/release-log/internal/app/git"

// Tag is the interface for the tag package
type Tag interface {
	List() ([]string, error)
	ListWithArgs(args []string) ([]string, error)
	SortedList(sortKey string) ([]string, error)
}

// Obj is the data structure struct for the tag package
type Obj struct {
	Git git.Git
}

// New returns a new intance of Tag interface
func New(g git.Git) Tag {
	return &Obj{Git: g}
}
