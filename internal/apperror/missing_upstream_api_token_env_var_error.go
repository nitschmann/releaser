package apperror

import "fmt"

type MissingUpstreamAPITokenEnvVarError struct {
	UpstreamName string
}

// NewMissingUpstreamAPITokenEnvVarError returns a new instance of MissingUpstreamAPITokenEnvVarError
func NewMissingUpstreamAPITokenEnvVarError(upstreamName string) *MissingUpstreamAPITokenEnvVarError {
	return &MissingUpstreamAPITokenEnvVarError{
		UpstreamName: upstreamName,
	}
}

// Error prints the actual error message
func (err *MissingUpstreamAPITokenEnvVarError) Error() string {
	return fmt.Sprintf("missing 'api_token_env_var' configuration for upstream '%s'", err.UpstreamName)
}
