package helper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nitschmann/releaser/internal/helper"
)

type commandContexSuite struct {
	suite.Suite
}

func TestCommandContextSuite(t *testing.T) {
	suite.Run(t, new(commandContexSuite))
}

func (s *commandContexSuite) TestNewContextWithCommandExecutionPath() {
	path := "~/code/releaser"
	ctx := helper.NewContextWithCommandExecutionPath(context.TODO(), path)

	s.Equal(helper.CommandExecutionPathFromContext(ctx), path)
}

func (s *commandContexSuite) TestCommandExecutionPathFromContext() {
	s.Equal(helper.CommandExecutionPathFromContext(context.TODO()), "")
}
