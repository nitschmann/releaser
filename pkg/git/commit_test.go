package git_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/test/mock"
)

type commitTestSuite struct {
	suite.Suite

	git mock.Git
}

func TestCommitSuite(t *testing.T) {
	suite.Run(t, new(commitTestSuite))
}

func (s *commitTestSuite) SetupTest() {
	s.git = mock.Git{}
}

func (s *commitTestSuite) TearDownTest() {
	s.git.AssertExpectations(s.T())
}

// func (s *commitTestSuite) TestLogsBetweenVersions() {
// 	s.Run("without versions", func() {
// 		commitLogs := []git.CommitLog{
// 			{
// 				Hash:      "85466a06e883a139ad72d1354a4307c",
// 				Message:   "Commit 2",
// 				Timestamp: strconv.Itoa(int(time.Now().Unix())),
// 			},
// 			{
// 				Hash:      "3d541938fb91e3e7ddb0bd2439a81ae23a",
// 				Message:   "Commit 1",
// 				Timestamp: strconv.Itoa(int(time.Now().Add(-24 * time.Hour).Unix())),
// 			},
// 		}

// 		var b bytes.Buffer

// 		for _, msg := range commitLogs {
// 			b.WriteString(fmt.Sprintf(
// 				`{"hash":"%s", "message":"%s", "timestamp":"%s"}`,
// 				msg.Hash,
// 				msg.Message,
// 				msg.Timestamp,
// 			) + "\n",
// 			)
// 		}

// 		gitMsg := fmt.Sprintf("%s\n\n\n", b.String())
// 		formatStr := `format:{"hash":"%H","message":"%s","timestamp":"%at"}`
// 		gitCmdArgs := []string{"log", "--oneline", fmt.Sprintf("--pretty='%s'", formatStr)}

// 		s.git.On("ExecCommand", gitCmdArgs).Once().Return(gitMsg, 0, nil)

// 		commit := git.NewCommit(s.git)

// 		logs, err := commit.LogsBetweenVersions("", "")

// 		s.NoError(err)
// 		s.Len(logs, 2)
// 		s.Equal(logs, commitLogs)
// 	})

// 	s.Run("with versions", func() {
// 		commitLogs := []git.CommitLog{
// 			{
// 				Hash:      "85466a06e883a139ad72d1354a4307c",
// 				Message:   "Commit 2",
// 				Timestamp: strconv.Itoa(int(time.Now().Unix())),
// 			},
// 			{
// 				Hash:      "3d541938fb91e3e7ddb0bd2439a81ae23a",
// 				Message:   "Commit 1",
// 				Timestamp: strconv.Itoa(int(time.Now().Add(-24 * time.Hour).Unix())),
// 			},
// 		}

// 		var b bytes.Buffer

// 		for _, msg := range commitLogs {
// 			b.WriteString(fmt.Sprintf(
// 				`{"hash":"%s", "message":"%s", "timestamp":"%s"}`,
// 				msg.Hash,
// 				msg.Message,
// 				msg.Timestamp,
// 			) + "\n",
// 			)
// 		}

// 		versionA := "v1.0.0"
// 		versionB := "master"
// 		gitMsg := fmt.Sprintf("%s\n\n\n", b.String())
// 		formatStr := `format:{"hash":"%H","message":"%s","timestamp":"%at"}`
// 		versionStr := fmt.Sprintf("%s..%s", versionA, versionB)
// 		gitCmdArgs := []string{"log", "--oneline", fmt.Sprintf("--pretty=%s", formatStr), versionStr}

// 		s.git.On("ExecCommand", gitCmdArgs).Once().Return(gitMsg, 0, nil)

// 		commit := git.NewCommit(s.git)

// 		logs, err := commit.LogsBetweenVersions(versionA, versionB)

// 		s.NoError(err)
// 		s.Len(logs, 2)
// 		s.Equal(logs, commitLogs)
// 	})

// }
