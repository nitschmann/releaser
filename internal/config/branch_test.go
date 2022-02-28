package config_test

import (
	"testing"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/stretchr/testify/suite"
)

type branchTestSuite struct {
	suite.Suite
}

func TestBranchSuite(t *testing.T) {
	suite.Run(t, new(branchTestSuite))
}

func (s *branchTestSuite) TestGetAllowedWithoutType() {
	val := config.Branch{}.GetAllowedWithoutType()
	s.Equal(val, config.BranchAllowedWithoutTypeDefault)
}

func (s *branchTestSuite) TestGetDelimiter() {
	val := config.Branch{}.GetDelimiter()
	s.Equal(val, config.BranchDelimiterDefault)
}

func (s *branchTestSuite) TestGetNameFormat() {
	val := config.Branch{}.GetNameFormat()
	s.Equal(val, config.BranchNameFormatDefault)
}

func (s *branchTestSuite) TestGetTypes() {
	val := config.Branch{}.GetTypes()
	s.Equal(val, config.BranchTypesDefault)
}
