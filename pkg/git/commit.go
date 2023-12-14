package git

import (
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
	// Author (name) of the commit
	Author string
	// AuthorEmail of the commit
	AuthorEmail string
	// Hash to identify the commit
	Hash string
	// Message of the commit
	Message   string
	Timestamp string
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

	formatStr := `format:'%H' + '%s' + '%at' + '%an' + '%ae'`
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

		gitCommitLog = c.removeStringSurroundingSingleQuotes(gitCommitLog)
		gitCommitLog = strings.TrimPrefix(gitCommitLog, "format:")

		gitCommitLogParts := strings.Split(gitCommitLog, " + ")
		for i, v := range gitCommitLogParts {
			gitCommitLogParts[i] = c.removeStringSurroundingSingleQuotes(v)
		}

		log.Hash = gitCommitLogParts[0]
		log.Message = gitCommitLogParts[1]
		log.Timestamp = gitCommitLogParts[2]
		log.Author = gitCommitLogParts[3]
		log.AuthorEmail = gitCommitLogParts[4]

		logs = append(logs, log)
	}

	return logs, nil
}

func (c commit) removeStringSurroundingSingleQuotes(str string) string {
	str = strings.TrimPrefix(str, "'")
	str = strings.TrimSuffix(str, "'")

	return str
}
