package mock

import (
	"github.com/stretchr/testify/mock"
)

type Git struct {
	mock.Mock
}

func (g Git) ExecCommand(args []string) (string, int, error) {
	calledArgs := g.Called(args)
	return calledArgs.Get(0).(string), calledArgs.Get(1).(int), calledArgs.Error(2)
}
