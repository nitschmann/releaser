package config

import (
	"reflect"
	"testing"
)

func TestRule_FlagsForBranch(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want []DynamicFlag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FlagsForBranch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.FlagsForBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_FlagsForCommit(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want []DynamicFlag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FlagsForCommit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.FlagsForCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_FlagNames(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FlagNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.FlagNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_FlagNamesForBranch(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FlagNamesForBranch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.FlagNamesForBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_FlagNamesForCommit(t *testing.T) {
	tests := []struct {
		name string
		r    Rule
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.FlagNamesForCommit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.FlagNamesForCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_MatchesWithPath(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		r       Rule
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MatchesWithPath(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rule.MatchesWithPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rule.MatchesWithPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_ParsedPaths(t *testing.T) {
	tests := []struct {
		name    string
		r       Rule
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ParsedPaths()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rule.ParsedPaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rule.ParsedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRule_Validate(t *testing.T) {
	tests := []struct {
		name    string
		r       Rule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Rule.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
