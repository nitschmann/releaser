package git

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Commit is the interface to abstract git commit handling
type Commit interface {
	// New creates a new git commit with the specified message
	New(message string) error
	// LogsBetweenVersions returns a log between two commit versions
	// If used with empty strings as versionA and B it returns the full current log
	LogsBetweenVersions(versionA, versionB string) ([]CommitLog, error)
}

// CommitLog is the data structure for a log of a git commit
type CommitLog struct {
	Hash      string `json:"hash"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
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

func (c commit) LogsBetweenVersions(versionA, versionB string) ([]CommitLog, error) {
	var logs []CommitLog

	formatStr := `format:{"hash":"%H","message":"%s","timestamp":"%at"}`
	gitCmdArgs := []string{"log", "--oneline", fmt.Sprintf("--pretty='%s'", formatStr)}

	if versionA != "" && versionB != "" {
		gitCmdArgs = append(gitCmdArgs, fmt.Sprintf("%s..%s", versionA, versionB))
	}

	gitCommitLogsStr, _, err := c.Git.ExecCommand(gitCmdArgs)
	if err != nil {
		return logs, err
	}

	gitCommitLogs := cleanEmptyEntriesFromStringSlice(strings.Split(gitCommitLogsStr, "\n"))
	for _, gitCommitLog := range gitCommitLogs {
		var log CommitLog

		gitCommitLog = strings.TrimPrefix(gitCommitLog, "'")
		gitCommitLog = strings.TrimPrefix(gitCommitLog, "format:")
		gitCommitLog = strings.TrimSuffix(gitCommitLog, "'")

		err = json.Unmarshal([]byte(gitCommitLog), &log)
		if err != nil {
			return logs, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}
