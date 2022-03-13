package service

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/internal/helper"
)

// PromptYesOrNoWithExpectedYes prompts the user in the CLI with 'yes' and 'no' and expects 'yes' to be choosen, else error.
func PromptYesOrNoWithExpectedYes(promptMsg string) error {
	fmt.Print(promptMsg)
	prompt := promptui.Select{
		Label: promptMsg,
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	if result != "Yes" {
		return apperror.NewPromptAbortError()
	}

	return nil
}

// ValidateCustomFlags checks if the specified map of flag matches with the valid flag names
func ValidateCustomFlags(validFlagNames []string, flags map[string]string) error {
	for flagName := range flags {
		if !helper.StringSliceIncludesElement(validFlagNames, flagName) {
			return apperror.NewInvalidFlagError(flagName)
		}
	}

	return nil
}

// ValidateType validate if the specified type value is part of the allowedTypes list
func ValidateType(allowedTypes []string, t string, allowEmpty bool) error {
	if allowEmpty && t == "" {
		return nil
	}

	if !helper.StringSliceIncludesElement(allowedTypes, t) {
		return apperror.NewInvalidFlagValueError("type", t)
	}

	return nil
}
