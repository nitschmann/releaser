package git

import (
	"os/exec"
	"strings"
)

// Git is the git package interface to abstract git command executions
type Git interface {
	ExecCommand(args []string) (string, int, error)
}

type gitObj struct {
	Executable string
}

// New returns a new pointer instance of Git interface with specified 'git' executable
func New(executable string) Git {
	return &gitObj{Executable: executable}
}

// ExecCommand executes a git command with the given args
func (g gitObj) ExecCommand(args []string) (string, int, error) {
	cmd := exec.Command(g.Executable, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		if ok {
			return "", exitError.ExitCode(), NewCommandError(err, cmd, output)

		} else {
			return "", -1, NewCommandError(err, cmd, output)
		}
	}

	return strings.TrimSuffix(string(output[:]), "\n"), 0, nil
}
