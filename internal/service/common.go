package service

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/nitschmann/releaser/internal/apperror"
)

// PromptYesOrNoWithExpectedYes prompts the user in the CLI with 'yes' and 'no' and expects 'yes' to be choosen, else error.
func PromptYesOrNoWithExpectedYes(promptMsg string) error {
	fmt.Print(promptMsg)
	prompt := promptui.Select{
		// Label: promptMsg,
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
