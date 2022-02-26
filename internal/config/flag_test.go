package config_test

import (
	"fmt"
	"testing"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/stretchr/testify/suite"
)

type flagTestSuite struct {
	suite.Suite
}

func TestFlagSuite(t *testing.T) {
	suite.Run(t, new(flagTestSuite))
}

func (s *flagTestSuite) TestGetDescription() {
	flag := config.Flag{Name: "flag1"}
	val := flag.GetDescription()
	s.Equal(val, fmt.Sprintf("Custom flag '%s'", flag.Name))
}

func (s *flagTestSuite) TestGetRequired() {
	val := config.Flag{}.GetRequired()
	s.Equal(val, config.FlagRequiredDefault)
}

func (s *flagTestSuite) TestGetSkipForBranch() {
	val := config.Flag{}.GetSkipForBranch()
	s.Equal(val, config.FlagSkipForBranchDefault)
}

func (s *flagTestSuite) TestGetSkipForCommit() {
	val := config.Flag{}.GetSkipForCommit()
	s.Equal(val, config.FlagSkipForCommitDefault)
}

func (s *flagTestSuite) TestGetSkipForRelease() {
	val := config.Flag{}.GetSkipForRelease()
	s.Equal(val, config.FlagSkipForReleaseDefault)
}
