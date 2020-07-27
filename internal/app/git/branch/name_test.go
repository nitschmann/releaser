package branch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	type args struct {
		delimiter string
	}

	tests := []struct {
		name string
		args args
		want *Name
	}{
		{
			name: "default",
			args: args{delimiter: "-"},
			want: &Name{Delimiter: "-"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewName(tt.args.delimiter), tt.want)
		})
	}
}

func TestName_FormatStringWithRegexAndDelimiter(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		n    *Name
		args args
		want string
	}{
		{
			name: "with space",
			n:    NewName("-"),
			args: args{str: "new branch"},
			want: "new-branch",
		},
		{
			name: "with invalid chars",
			n:    NewName("+"),
			args: args{str: "super*feature*branch"},
			want: "super+feature+branch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.n.FormatStringWithRegexAndDelimiter(tt.args.str), tt.want)
		})
	}
}

func TestName_Join(t *testing.T) {
	tests := []struct {
		name string
		n    Name
		want string
	}{
		{
			name: "without prefix given",
			n: Name{
				Delimiter: "-",
				Suffix:    "branch",
				Title:     "new",
			},
			want: "new-branch",
		},
		{
			name: "without prefix given",
			n: Name{
				Delimiter: "-",
				Suffix:    "branch",
				Title:     "new",
			},
			want: "new-branch",
		},
		{
			name: "with prefix given",
			n: Name{
				Delimiter: "_",
				Prefix:    "ticket-123",
				Suffix:    "branch",
				Title:     "new_feature",
			},
			want: "ticket-123_new_feature_branch",
		},
		{
			name: "without suffix given",
			n: Name{
				Delimiter: "_",
				Prefix:    "ticket-123",
				Suffix:    "",
				Title:     "new_feature",
			},
			want: "ticket-123_new_feature",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.n.Join(), tt.want)
		})
	}
}

// func TestName_SetPrefixWithTemplate(t *testing.T) {
// 	type args struct {
// 		prefixTemplatePattern string
// 		templateValues        map[string]string
// 	}
// 	tests := []struct {
// 		name    string
// 		n       *Name
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.n.SetPrefixWithTemplate(tt.args.prefixTemplatePattern, tt.args.templateValues); (err != nil) != tt.wantErr {
// 				t.Errorf("Name.SetPrefixWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestName_SetSuffixWithTemplate(t *testing.T) {
// 	type args struct {
// 		suffixTemplatePattern string
// 		templateValues        map[string]string
// 	}
// 	tests := []struct {
// 		name    string
// 		n       *Name
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.n.SetSuffixWithTemplate(tt.args.suffixTemplatePattern, tt.args.templateValues); (err != nil) != tt.wantErr {
// 				t.Errorf("Name.SetSuffixWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestName_SetTitleWithTemplate(t *testing.T) {
// 	type args struct {
// 		titleTemplatePattern string
// 		templateValues       map[string]string
// 	}
// 	tests := []struct {
// 		name    string
// 		n       *Name
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.n.SetTitleWithTemplate(tt.args.titleTemplatePattern, tt.args.templateValues); (err != nil) != tt.wantErr {
// 				t.Errorf("Name.SetTitleWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestName_ValidCharsRegex(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		n    Name
// 		want *regexp.Regexp
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.n.ValidCharsRegex(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Name.ValidCharsRegex() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
