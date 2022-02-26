package cmd

import (
	"fmt"
	"os"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/pkg/git"
)

func printCLIErrorAndExit(err error) {
	switch e := err.(type) {
	case *apperror.ConfigValidationErrors:
		fmt.Println(err.Error())
	case *apperror.PromptAbortError:
		fmt.Printf("%s\n", err.Error())
	case *git.CommandError:
		fmt.Printf("git execution error: \n%s\n", e.Output)
	default:
		fmt.Printf("An unexpected error occurred:\n%v\n", err.Error())
	}

	os.Exit(1)
}
