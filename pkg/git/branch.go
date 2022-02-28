package git

// Branch is the interface to abstract git branch handling
type Branch interface {
	// Checkout a new branch with the specified name
	Checkout(name string) error
	// Current returns the name of the current branch
	Current() (string, error)
	// ValidateName checks if a given name is a potentially valid branch name
	ValidateName(name string) error
}

type branch struct {
	Git Git
}

// NewBranch returns an instance of Branch interface with default values
func NewBranch(gitObj Git) Branch {
	return branch{Git: gitObj}
}

func (b branch) Checkout(name string) error {
	_, _, err := b.Git.ExecCommand([]string{"checkout", "-b", name})
	return err
}

func (b branch) Current() (string, error) {
	branchName, _, err := b.Git.ExecCommand([]string{"branch", "--show-current"})
	if err != nil {
		return "", err
	}

	return branchName, nil
}

func (b branch) ValidateName(name string) error {
	_, _, err := b.Git.ExecCommand([]string{"check-ref-format", "--branch", name})

	if err != nil {
		return NewInvalidBranchNameError(err, name)
	}

	return nil
}
