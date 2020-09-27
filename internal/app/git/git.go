package git

import (
	"os/exec"
	"strings"
)

// Git is the git package interface
type Git interface {
	ExecCommand(args []string) (string, error)
}

// Obj is the git package data struct
type Obj struct {
	Executable string
}

// New returns a new pointer instance of Git
func New(executable string) Git {
	return &Obj{Executable: executable}
}

// ExecCommand executes a git command with the given args
func (g Obj) ExecCommand(args []string) (string, error) {
	cmd := exec.Command(g.Executable, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", CmdError{
			Cmd:     cmd,
			Err:     err,
			Message: string(output),
		}
	}

	return strings.TrimSuffix(string(output[:]), "\n"), nil
}
