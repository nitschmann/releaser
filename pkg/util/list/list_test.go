package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanEmptyStrings(t *testing.T) {
	type args struct {
		list []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "with empty list",
			args: args{list: []string{}},
			want: []string{},
		},
		{
			name: "without empty strings",
			args: args{list: []string{"value 1", "value 2"}},
			want: []string{"value 1", "value 2"},
		},
		{
			name: "with single empty string included",
			args: args{list: []string{"value", ""}},
			want: []string{"value"},
		},
		{
			name: "with multiple empty strings included",
			args: args{list: []string{"value", "", "", "value 2"}},
			want: []string{"value", "value 2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, CleanEmptyStrings(tt.args.list), tt.want)
		})
	}
}
