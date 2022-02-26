package apperror

type PromptAbortError struct{}

// NewPromptAbortError returns a new instance of PromptAbortError
func NewPromptAbortError() *PromptAbortError {
	return &PromptAbortError{}
}

// Error prints the actual error message
func (err *PromptAbortError) Error() string {
	return "execution aborted with prompt input"
}
