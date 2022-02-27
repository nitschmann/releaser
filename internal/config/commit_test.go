package config_test

import (
	"testing"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/stretchr/testify/suite"
)

type commitTestSuite struct {
	suite.Suite
}

func TestCommitSuite(t *testing.T) {
	suite.Run(t, new(commitTestSuite))
}

func (s *commitTestSuite) TestGetAllowedWithoutType() {
	val := config.Commit{}.GetAllowedWithoutType()
	s.Equal(val, config.CommitAllowedWithoutTypeDefault)
}

func (s *commitTestSuite) TestGetMessageFormat() {
	val := config.Commit{}.GetMessageFormat()
	s.Equal(val, config.CommitMessageFormatDefault)
}

func (s *commitTestSuite) TestGetTypes() {
	val := config.Commit{}.GetTypes()
	s.Equal(val, config.CommitTypesDefault)
}
