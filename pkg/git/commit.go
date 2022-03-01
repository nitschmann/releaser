package git

// Commit is the interface to abstract git commit handling
type Commit interface {
	// New creates a new git commit with the specified message
	New(message string) error
}

type commit struct {
	Git Git
}

// NewCommit returns an instance of Commit interface with default values
func NewCommit(gitObj Git) Commit {
	return commit{Git: gitObj}
}

func (c commit) New(message string) error {
	_, _, err := c.Git.ExecCommand([]string{"commit", "-m", message})
	if err != nil {
		return err
	}

	return nil
}
