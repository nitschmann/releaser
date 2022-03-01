package apperror

import "fmt"

// InvalidFlagError if a custom flag is invalid
type InvalidFlagError struct {
	Flag string
}

// NewInvalidFlagError returns a new instance of InvalidFlagError
func NewInvalidFlagError(flag string) *InvalidFlagError {
	return &InvalidFlagError{
		Flag: flag,
	}
}

// Error prints the actual error message
func (err *InvalidFlagError) Error() string {
	return fmt.Sprintf("invalid or unknown flag '%s'", err.Flag)
}
