package apperror

import "fmt"

// InvalidUpstreamNameError if a upstream name is invalid
type InvalidUpstreamNameError struct {
	Name       string
	ValidNames []string
}

// NewInvalidUpstreamNameError returns a new instance of InvalidUpstreamNameError
func NewInvalidUpstreamNameError(name string, validNames []string) *InvalidUpstreamNameError {
	return &InvalidUpstreamNameError{
		Name:       name,
		ValidNames: validNames,
	}
}

// Error prints the actual error message
func (err *InvalidUpstreamNameError) Error() string {
	return fmt.Sprintf("Invalid upstream '%s' given. Valid upstreams are: %v", err.Name, err.ValidNames)
}
