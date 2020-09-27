package config

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/go-playground/validator/v10"
	"github.com/mattn/go-zglob"
	"github.com/spf13/viper"
)

// Rule is a struct of settings and rules for a certain path
type Rule struct {
	// A list of dynamic flags which are added to certain commands under the given paths.
	Flags []DynamicFlag `mapstructure:"flags" yaml:"flags"`
	// GitBranchDefaultType  defines the default type of branch that should be used
	GitBranchDefaultType string `mapstructure:"git_branch_default_type" yaml:"git_branch_default_type"`
	// GitBranchDelimiter defines the sign which connects certain parts of the branch name.
	// It just can be one sign and only one of '-', '_', '/' '+'. Default is '-'
	GitBranchDelimiter string `mapstructure:"git_branch_delimiter" validate:"omitempty,max=1,oneof='-' '_' '/' '+'" yaml:"git_branch_delimiter"`
	// GitBranchForceType: if true a type for the branch needs always to be set
	GitBranchForceType bool `mapstructure:"git_branch_force_type" yaml:"git_branch_force_type"`
	// GitBranchPrefix defines the prefix which is set for a branch name
	GitBranchPrefix string `mapstructure:"git_branch_prefix" yaml:"git_branch_prefix"`
	// GitBranchSuffix defines the suffix which is set for a branch name
	GitBranchSuffix string `mapstructure:"git_branch_suffix" yaml:"git_branch_suffix"`
	// GitBranchTypes defines the overwrites of available types. If not set the defaults are:
	// feature, bug, hotfix
	GitBranchTypes []string `mapstructure:"git_branch_types" validate:"unique,dive,lowercase,alpha" yaml:"git_branch_types"`
	// Paths are the paths with certain rules to which this rule should be applied to
	Paths []string `mapstructure:"paths" yaml:"paths"`
}

// FlagsForBranch will only return entries of r.Flags list where SkipForBranch is false
func (r Rule) FlagsForBranch() []DynamicFlag {
	var list []DynamicFlag

	for _, flag := range r.Flags {
		if !flag.SkipForBranch {
			list = append(list, flag)
		}
	}

	return list
}

// FlagsForCommit will only return entries of r.Flags list where SkipForCommit is false
func (r Rule) FlagsForCommit() []DynamicFlag {
	var list []DynamicFlag

	for _, flag := range r.Flags {
		if !flag.SkipForCommit {
			list = append(list, flag)
		}
	}

	return list
}

// FlagNames returns the Name attribute of each element in r.Flags
func (r Rule) FlagNames() []string {
	var list []string

	for _, flag := range r.Flags {
		list = append(list, flag.Name)
	}

	return list
}

// FlagNamesForBranch returns the Name attribute of each element in r.FlagNamesForBranch
func (r Rule) FlagNamesForBranch() []string {
	var list []string

	for _, flag := range r.FlagsForBranch() {
		list = append(list, flag.Name)
	}

	return list
}

// FlagNamesForCommit returns the Name attribute of each element in r.FlagNamesForCommit
func (r Rule) FlagNamesForCommit() []string {
	var list []string

	for _, flag := range r.FlagsForCommit() {
		list = append(list, flag.Name)
	}

	return list
}

// GetGitBranchDelimiter returns the defined r.GitBranchDelimiter or "-" as default
func (r Rule) GetGitBranchDelimiter() string {
	if r.GitBranchDelimiter == "" {
		return "-"
	}

	return r.GitBranchDelimiter
}

// GetGitBranchTypes returns either r.GitBranchTypes or the default ones
func (r Rule) GetGitBranchTypes() []string {
	if len(r.GitBranchTypes) > 0 {
		return r.GitBranchTypes
	}

	return []string{"bug", "feature", "hotfix"}
}

// MatchesWithPath checks if a given path matches with any of the defined path rules
func (r Rule) MatchesWithPath(p string) (bool, error) {
	parsedPaths, err := r.ParsedPaths()
	if err != nil {
		return false, err
	}

	for _, parsedPathPattern := range parsedPaths {
		matched, err := zglob.Match(parsedPathPattern, p)
		if err != nil {
			return false, err
		}

		if !matched {
			continue
		}

		return matched, nil
	}

	return false, nil
}

// ParsedPaths returns the Paths list with parsed variables of the Rule instance
func (r Rule) ParsedPaths() ([]string, error) {
	var list []string

	for _, pathPatternRule := range r.Paths {
		pathPatternRuleTemplate, err := template.New("pathRule").Parse(pathPatternRule)
		if err != nil {
			return list, err
		}

		rules, err := r.pathRuleVariables()
		if err != nil {
			return list, err
		}

		buf := &bytes.Buffer{}
		err = pathPatternRuleTemplate.Execute(buf, rules)
		if err != nil {
			return list, err
		}

		list = append(list, buf.String())
	}

	return list, nil
}

func (r Rule) pathRuleVariables() (map[string]string, error) {
	rules := map[string]string{
		"ConfigParentDir": "",
		"UserHomeDir":     "",
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return rules, err
	}

	rules["ConfigParentDir"] = filepath.Dir(filepath.Dir(viper.ConfigFileUsed()))
	rules["ProjectDir"] = rules["ConfigParentDir"]
	rules["UserHomeDir"] = userHomeDir

	return rules, nil
}

// Validate runs the validators of the Rule instance
func (r Rule) Validate() error {
	v := validator.New()
	err := v.Struct(r)
	if err != nil {
		return err
	}

	err = r.validateFlags()
	if err != nil {
		return err
	}

	return nil
}

func (r Rule) validateFlags() error {
	for _, flag := range r.Flags {
		err := flag.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
