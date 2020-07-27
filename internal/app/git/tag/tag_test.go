package tag

import (
	"testing"

	"github.com/nitschmann/release-log/internal/app/git"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		g git.Git
	}

	gitObj := git.New("git")

	tests := []struct {
		name string
		args args
		want Tag
	}{
		{
			name: "default",
			args: args{g: gitObj},
			want: &Obj{Git: gitObj},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, New(tt.args.g), tt.want)
		})
	}
}
