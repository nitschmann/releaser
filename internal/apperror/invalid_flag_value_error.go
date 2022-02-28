package apperror

import "fmt"

type InvalidFlagValueError struct {
	Flag  string
	Value string
}

// NewInvalidFlagValueError returns a new instance of InvalidFlagValueError
func NewInvalidFlagValueError(flag, value string) *InvalidFlagValueError {
	return &InvalidFlagValueError{
		Flag:  flag,
		Value: value,
	}
}

// Error prints the actual error message
func (err *InvalidFlagValueError) Error() string {
	return fmt.Sprintf("value '%s' is invalid for flag '%s'", err.Value, err.Flag)
}
