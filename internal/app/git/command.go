package git

import (
	"os/exec"
	"strings"

	"github.com.com/nitschmann/release-log/internal/app/config"
)

// Executes a git command with the given args
func ExecCommand(args []string) (string, error) {
	output, err := exec.Command(config.Get().GitExecutable, args...).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output[:]), "\n"), nil
}
