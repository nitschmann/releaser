package git

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testGitExecCommandFailureArgs           []string = []string{"invalid-command"}
	testGitExecCommandFailureReturnErrorMsg string   = "git: 'invalid-command' is not a git command."
	testGitExecCommandSuccessArgs           []string = []string{"--version"}
	testGitExecCommandSuccessReturn         string   = "git version 2.24.3"
)

type testFailureGitMock struct {
	Obj
}

type testSuccessGitMock struct {
	Obj
}

func (g testFailureGitMock) ExecCommand(args []string) (string, error) {
	return "", errors.New(testGitExecCommandFailureReturnErrorMsg)
}

func (g testSuccessGitMock) ExecCommand(args []string) (string, error) {
	return testGitExecCommandSuccessReturn, nil
}

func TestNew(t *testing.T) {
	type args struct {
		executable string
	}

	tests := []struct {
		name string
		args args
		want Git
	}{
		{
			name: "with git executable",
			args: args{executable: "git"},
			want: &Obj{Executable: "git"},
		},
		{
			name: "with empty git executable string",
			args: args{executable: ""},
			want: &Obj{Executable: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, New(tt.args.executable), tt.want)
		})
	}
}

func TestObj_ExecCommand(t *testing.T) {
	type execCommandFunc func(args []string) (string, error)

	type args struct {
		args []string
	}

	tests := []struct {
		name      string
		git       Git
		args      args
		want      string
		wantErr   bool
		errString string
	}{
		{
			name:    "successful command",
			git:     &testSuccessGitMock{},
			args:    args{args: testGitExecCommandSuccessArgs},
			want:    testGitExecCommandSuccessReturn,
			wantErr: false,
		},
		{
			name:      "failure command",
			git:       &testFailureGitMock{},
			args:      args{args: testGitExecCommandFailureArgs},
			want:      "",
			wantErr:   true,
			errString: testGitExecCommandFailureReturnErrorMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.git.ExecCommand(tt.args.args)
			assert.Equal(t, result, tt.want)
			if tt.wantErr {
				assert.Error(t, err, tt.errString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
