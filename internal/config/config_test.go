package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/helper"
)

type configTestSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(configTestSuite))
}

func (s *configTestSuite) TestValidate() {
	s.Run("with default values", func() {
		c := config.New()
		err := c.Validate()

		s.NoError(err)
	})

	s.Run("Branch", func() {
		s.Run("Delimiter", func() {
			s.Run("too long", func() {
				c := config.New()
				c.Branch.Delimiter = helper.StringToPointer("too-long")
				err := c.Validate()
				s.Error(err)

				validationErr := err.(*apperror.ConfigValidationErrors)
				s.Equal(len(validationErr.Errors), 1)
				s.Equal(validationErr.Errors[0].Error(), "branch.delimiter must be 1 character in length")
			})
		})

		s.Run("Types", func() {
			s.Run("entry with non-alphanumeric characters", func() {
				c := config.New()
				c.Branch.Types = []string{"invalid-type"}
				err := c.Validate()
				s.Error(err)

				validationErr := err.(*apperror.ConfigValidationErrors)
				s.Equal(len(validationErr.Errors), 1)
				s.Equal(validationErr.Errors[0].Error(), "branch.types[0] can only contain alphanumeric characters")
			})
			s.Run("not all entries unique", func() {
				c := config.New()
				c.Branch.Types = []string{"fix", "fix"}
				err := c.Validate()
				s.Error(err)

				validationErr := err.(*apperror.ConfigValidationErrors)
				s.Equal(len(validationErr.Errors), 1)
				s.Equal(validationErr.Errors[0].Error(), "branch.types must contain unique values")
			})
		})
	})

	s.Run("Commit", func() {
		s.Run("Types", func() {
			s.Run("entry with non-alphanumeric characters", func() {
				c := config.New()
				c.Commit.Types = []string{"invalid-type"}
				err := c.Validate()
				s.Error(err)

				validationErr := err.(*apperror.ConfigValidationErrors)
				s.Equal(len(validationErr.Errors), 1)
				s.Equal(validationErr.Errors[0].Error(), "commit.types[0] can only contain alphanumeric characters")
			})
			s.Run("not all entries unique", func() {
				c := config.New()
				c.Commit.Types = []string{"fix", "fix"}
				err := c.Validate()
				s.Error(err)

				validationErr := err.(*apperror.ConfigValidationErrors)
				s.Equal(len(validationErr.Errors), 1)
				s.Equal(validationErr.Errors[0].Error(), "commit.types must contain unique values")
			})

		})
	})
}
