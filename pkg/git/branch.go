package git

// Branch is the interface to abstract git branch handling
type Branch interface {
	ValidateName(name string) error
}

type branch struct {
	Git Git
}

// NewBranch returns an instance of Branch interface with default values
func NewBranch(gitObj Git) Branch {
	return branch{Git: gitObj}
}

// ValidateName checks if a given name is a potentially valid branch name
func (b branch) ValidateName(name string) error {
	_, _, err := b.Git.ExecCommand([]string{"check-ref-format", "--branch", name})

	if err != nil {
		return NewInvalidBranchNameError(err, name)
	}

	return nil
}
