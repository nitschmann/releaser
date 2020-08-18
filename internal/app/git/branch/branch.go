package branch

import "github.com/nitschmann/release-log/internal/app/git"

// Branch is the interface for the branch package
type Branch interface {
	Checkout(branchName string) error
}

// Obj is the data structure struct for the branch package
type Obj struct {
	Git git.Git
}

// New returns a new intance of Tag interface
func New(g git.Git) Branch {
	return &Obj{Git: g}
}
