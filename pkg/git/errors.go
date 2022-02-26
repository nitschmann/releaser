package git

import (
	"fmt"
	"os/exec"
)

// CommandError in case a git command execution fails
type CommandError struct {
	Err error
	// The actually executed OS command
	Cmd *exec.Cmd
	// Human readable string of the command output
	Output string
}

// NewCommandError returns an new instance of CommandError with specified values
func NewCommandError(err error, cmd *exec.Cmd, cmdOutput []byte) *CommandError {
	return &CommandError{
		Cmd:    cmd,
		Err:    err,
		Output: string(cmdOutput),
	}
}

// Error returns the actual error message
func (err CommandError) Error() string {
	return err.Output
}

// InvalidBranchNameError is returned if a branch has a invalid name
type InvalidBranchNameError struct {
	Err  error
	Name string
}

// NewInvalidBranchNameError returns an new instance of InvalidBranchNameError
func NewInvalidBranchNameError(err error, name string) *InvalidBranchNameError {
	return &InvalidBranchNameError{
		Err:  err,
		Name: name,
	}
}

// Error returns the actual error message
func (err InvalidBranchNameError) Error() string {
	return fmt.Sprintf("Branch '%s' is not a valid branch name", err.Name)
}
