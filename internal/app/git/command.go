package git

import (
	"os/exec"
	"strings"

	"github.com/nitschmann/release-log/internal/app/config"
)

// ExecCommand executes a git command with the given args
// TODO: Pass the git executable directly in and don't use config here
func ExecCommand(args []string) (string, error) {
	output, err := exec.Command(config.Get().GitExecutable, args...).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output[:]), "\n"), nil
}
