package helper

import "context"

type commandContextKey int

var commandExecutionPathKey = commandContextKey(1)

// NewContextWithCommandExecutionPath sets the given path argument as value into a new context.Context
func NewContextWithCommandExecutionPath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, commandExecutionPathKey, path)
}

// CommandExecutionPathFromContext returns the commandExecutionPath in ctx, if any
func CommandExecutionPathFromContext(ctx context.Context) string {
	commandExecutionPath, _ := ctx.Value(commandExecutionPathKey).(string)
	return commandExecutionPath
}
