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
	output, err := exec.Command(g.Executable, args...).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output[:]), "\n"), nil
}
