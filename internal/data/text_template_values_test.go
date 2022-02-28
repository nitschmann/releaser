package data_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/data"
)

type textTemplateValuesTestSuite struct {
	suite.Suite
}

func TestTextTemplateValuesSuite(t *testing.T) {
	suite.Run(t, new(textTemplateValuesTestSuite))
}

func (s *textTemplateValuesTestSuite) TestAdd() {
	key := "key1"
	value := "value 1"
	templateValues := data.NewTextTemplateValues()
	templateValues.AddFlag(key, value)

	s.Equal(templateValues.Flags[key], value)
}

func (s *textTemplateValuesTestSuite) TestParseTemplateString() {
	s.Run("with given field", func() {
		branchType := "feature"
		templateValues := data.NewTextTemplateValues()
		templateValues.Type = branchType
		templateString := "{{ .Type }}-branch"

		result, err := templateValues.ParseTemplateString(templateString)

		s.NoError(err)
		s.Equal(result, fmt.Sprintf("%s-branch", branchType))
	})

	s.Run("with empty field being used", func() {
		templateValues := data.NewTextTemplateValues()
		templateString := "{{ if .Type }}{{ .Type }}-{{ end }}branch"

		result, err := templateValues.ParseTemplateString(templateString)

		s.NoError(err)
		s.Equal(result, "branch")
	})
}
