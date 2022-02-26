package apperror

import (
	"fmt"
	"strings"

	validatorPkg "github.com/go-playground/validator/v10"

	"github.com/nitschmann/releaser/internal/helper"
)

// ConfigValidationError is the error in case of any validation problems in the config file
type ConfigValidationError struct {
	Err validatorPkg.FieldError
	Msg string
}

// NewConfigValidationError returns an instance of ConfigValidationError
func NewConfigValidationError(err validatorPkg.FieldError, translatedErrorMsg string) *ConfigValidationError {
	namespaceParts := strings.Split(err.Namespace(), ".")
	namespaceParts = helper.RemoveElementFromStringSlice(namespaceParts, 0)
	namespaceParts = helper.RemoveElementFromStringSlice(namespaceParts, len(namespaceParts)-1)
	msg := fmt.Sprintf("%s.%s", strings.Join(namespaceParts, "."), translatedErrorMsg)

	return &ConfigValidationError{
		Err: err,
		Msg: msg,
	}
}

// Error prints the actual error message
func (e *ConfigValidationError) Error() string {
	return e.Msg
}

// ConfigValidationErrors is a list of multiple ConfigValidationError
type ConfigValidationErrors struct {
	Errors []*ConfigValidationError
}

// NewConfigValidationErrors returns an instance of ConfigValidationErrors
func NewConfigValidationErrors() *ConfigValidationErrors {
	return new(ConfigValidationErrors)
}

// Add a new ConfigValidationError to the Errors field
func (errs *ConfigValidationErrors) Add(err *ConfigValidationError) {
	errs.Errors = append(errs.Errors, err)
}

// Error prints the actual error message
func (errs *ConfigValidationErrors) Error() string {
	if !errs.HasErrors() {
		return ""
	}

	var sb strings.Builder

	sb.WriteString("Validation of releaser config failed:\n\n")

	for _, err := range errs.Errors {
		sb.WriteString(err.Error() + "\n")
	}

	return sb.String()
}

// HasErrors checks if there are any entries in Errors field
func (errs *ConfigValidationErrors) HasErrors() bool {
	return len(errs.Errors) > 0
}
