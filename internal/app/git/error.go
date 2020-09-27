package git

import "os/exec"

// CmdError reports an unsuccessful execution of a git command with exec.Cmd
type CmdError struct {
	Cmd     *exec.Cmd
	Err     error
	Message string
}

// Error returns the actual Message of CmdError instance
func (e CmdError) Error() string {
	return e.Message
}
