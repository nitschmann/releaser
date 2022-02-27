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

func (s *branchTestSuite) TestGetTitleFormat() {
	val := config.Branch{}.GetTitleFormat()
	s.Equal(val, config.BranchTitleFormatDefault)
}

func (s *branchTestSuite) TestGetTypes() {
	val := config.Branch{}.GetTypes()
	s.Equal(val, config.BranchTypesDefault)
}
