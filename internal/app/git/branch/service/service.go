package service

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nitschmann/releaser/internal/app/config"
	"github.com/nitschmann/releaser/internal/app/git"
	"github.com/nitschmann/releaser/internal/app/git/branch"
)

// Service a the branch service struct
type Service struct {
	ConfigRule     config.Rule
	GitService     git.Git
	TemplateValues map[string]string
}

func addTransformationsValuesToTemplateValues(templateValues map[string]string) map[string]string {
	newTemplateValues := make(map[string]string)

	for key, value := range templateValues {
		newTemplateValues[key] = value
		newTemplateValues[key+"Title"] = strings.Title(value)
		newTemplateValues[key+"Upper"] = strings.ToUpper(value)
	}

	return newTemplateValues
}

// New returns a new pointer instance of Service with the given arguments
func New(configRule config.Rule, gitService git.Git, templateValues map[string]string) *Service {
	return &Service{
		ConfigRule:     configRule,
		GitService:     gitService,
		TemplateValues: addTransformationsValuesToTemplateValues(templateValues),
	}
}

// BuildNewName builds a new name for a branch with the given title argument (can be also an empty string)
func (s Service) BuildNewName(branchType string, title string) (string, error) {
	var resultName string

	err := s.validateBranchType(branchType)
	if err != nil {
		return resultName, errors.New("invalid type given")
	}

	name := branch.NewName(s.ConfigRule.GitBranchDelimiter)

	err = name.SetPrefixWithTemplate(s.ConfigRule.GitBranchPrefix, s.TemplateValues)
	if err != nil {
		return resultName, err
	}

	err = name.SetTitleWithTemplate(title, s.TemplateValues)
	if err != nil {
		return resultName, err
	}

	err = name.SetSuffixWithTemplate(s.ConfigRule.GitBranchSuffix, s.TemplateValues)
	if err != nil {
		return resultName, err
	}

	resultName = name.Join()

	return resultName, nil
}

// Checkout tries to checkout a new branch with the given name
func (s Service) Checkout(branchName string) error {
	gitBranch := branch.New(s.GitService)
	return gitBranch.Checkout(branchName)
}

func (s Service) validateBranchType(t string) error {
	typesString := strings.Join(s.ConfigRule.GetGitBranchTypes(), " ")
	return validator.New().Var(t, "oneof="+typesString)
}
