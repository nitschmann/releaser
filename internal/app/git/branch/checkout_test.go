package branch

import (
	"errors"
	"testing"

	"github.com/nitschmann/releaser/internal/app/git"
	"github.com/stretchr/testify/assert"
)

var (
	testCheckoutGitExecCommandErrorReturn  error  = nil
	testCheckoutGitExecCommandResultReturn string = ""
)

type testCheckoutGitMock struct {
	git.Obj
}

func (g testCheckoutGitMock) ExecCommand(args []string) (string, error) {
	return testCheckoutGitExecCommandResultReturn, testCheckoutGitExecCommandErrorReturn
}

func TestObj_Checkout(t *testing.T) {
	type args struct {
		branchName string
	}

	errorMsg := "fatal: 'invalid branch' is not a valid branch name."
	tests := []struct {
		name            string
		branch          Branch
		args            args
		gitExecCmdError error
		wantErr         bool
		errString       string
	}{
		{
			name:            "default",
			branch:          New(&testCheckoutGitMock{}),
			args:            args{branchName: "new-branch"},
			gitExecCmdError: nil,
			wantErr:         false,
		},
		{
			name:            "with error",
			branch:          New(&testCheckoutGitMock{}),
			args:            args{branchName: "invalid branch"},
			gitExecCmdError: errors.New(errorMsg),
			wantErr:         true,
			errString:       errorMsg,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCheckoutGitExecCommandErrorReturn = tt.gitExecCmdError

			err := tt.branch.Checkout(tt.args.branchName)
			if tt.wantErr {
				assert.Error(t, err, tt.errString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
