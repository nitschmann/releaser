package tag

import (
	"errors"
	"strings"
	"testing"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/stretchr/testify/assert"
)

var (
	testListGitExecCommandResultReturn string = ""
	testListGitExecCommandErrorReturn  error  = nil
)

type testListGitMock struct {
	git.Obj
}

func (g testListGitMock) ExecCommand(args []string) (string, error) {
	return testListGitExecCommandResultReturn, testListGitExecCommandErrorReturn
}

func TestObj_List(t *testing.T) {
	var emptyTagList []string
	validTagList := []string{"v0.0.1", "v0.0.2", "v0.0.3"}
	errorMsg := "Couldn't execute git command"

	tests := []struct {
		name             string
		tag              Tag
		gitExecCmdResult string
		gitExecCmdError  error
		want             []string
		wantErr          bool
		errString        string
	}{
		{
			name:             "default",
			tag:              New(&testListGitMock{}),
			gitExecCmdResult: strings.Join(validTagList, "\n"),
			gitExecCmdError:  nil,
			want:             validTagList,
			wantErr:          false,
		},
		{
			name:             "with empty tag list",
			tag:              New(&testListGitMock{}),
			gitExecCmdResult: "",
			gitExecCmdError:  nil,
			want:             []string{},
			wantErr:          false,
		},
		{
			name:             "with error",
			tag:              New(&testListGitMock{}),
			gitExecCmdResult: "",
			gitExecCmdError:  errors.New(errorMsg),
			want:             emptyTagList,
			wantErr:          true,
			errString:        errorMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testListGitExecCommandResultReturn = tt.gitExecCmdResult
			testListGitExecCommandErrorReturn = tt.gitExecCmdError

			result, err := tt.tag.List()

			assert.Equal(t, result, tt.want)
			if tt.wantErr {
				assert.Error(t, err, tt.errString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestObj_ListWithArgs(t *testing.T) {
	type args struct {
		args []string
	}

	var emptyTagList []string
	validTagList := []string{"v0.0.1", "v0.0.2", "v0.0.3"}
	errorMsg := "error: unknown option `invalid-option`"

	tests := []struct {
		name             string
		tag              Tag
		args             args
		gitExecCmdResult string
		gitExecCmdError  error
		want             []string
		wantErr          bool
		errString        string
	}{
		{
			name:             "default",
			tag:              New(&testListGitMock{}),
			args:             args{args: []string{"--no-column"}},
			gitExecCmdResult: strings.Join(validTagList, "\n"),
			gitExecCmdError:  nil,
			want:             validTagList,
			wantErr:          false,
		},
		{
			name:             "with empty tag list",
			tag:              New(&testListGitMock{}),
			args:             args{args: []string{"--no-column"}},
			gitExecCmdResult: "",
			gitExecCmdError:  nil,
			want:             []string{},
			wantErr:          false,
		},
		{
			name:             "with error",
			tag:              New(&testListGitMock{}),
			args:             args{args: []string{"--invalid-option"}},
			gitExecCmdResult: "",
			gitExecCmdError:  errors.New(errorMsg),
			want:             emptyTagList,
			wantErr:          true,
			errString:        errorMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testListGitExecCommandResultReturn = tt.gitExecCmdResult
			testListGitExecCommandErrorReturn = tt.gitExecCmdError

			result, err := tt.tag.ListWithArgs(tt.args.args)

			assert.Equal(t, result, tt.want)
			if tt.wantErr {
				assert.Error(t, err, tt.errString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestObj_SortedList(t *testing.T) {
	type args struct {
		sortKey string
	}

	var emptyTagList []string
	validTagList := []string{"v0.0.1", "v0.0.2", "v0.0.3"}
	errorMsg := "unknown value for flag --sort"

	tests := []struct {
		name             string
		tag              Tag
		args             args
		gitExecCmdResult string
		gitExecCmdError  error
		want             []string
		wantErr          bool
		errString        string
	}{
		{
			name:             "default",
			tag:              New(&testListGitMock{}),
			args:             args{sortKey: "v:refname"},
			gitExecCmdResult: strings.Join(validTagList, "\n"),
			gitExecCmdError:  nil,
			want:             validTagList,
			wantErr:          false,
		},
		{
			name:             "with empty tag list",
			tag:              New(&testListGitMock{}),
			args:             args{sortKey: "v:refname"},
			gitExecCmdResult: "",
			gitExecCmdError:  nil,
			want:             []string{},
			wantErr:          false,
		},
		{
			name:             "with error",
			tag:              New(&testListGitMock{}),
			args:             args{sortKey: "invalid"},
			gitExecCmdResult: "",
			gitExecCmdError:  errors.New(errorMsg),
			want:             emptyTagList,
			wantErr:          true,
			errString:        errorMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testListGitExecCommandResultReturn = tt.gitExecCmdResult
			testListGitExecCommandErrorReturn = tt.gitExecCmdError

			result, err := tt.tag.SortedList(tt.args.sortKey)

			assert.Equal(t, result, tt.want)
			if tt.wantErr {
				assert.Error(t, err, tt.errString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
