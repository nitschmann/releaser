package project

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/manifoldco/promptui"

	"github.com/nitschmann/releaser/internal/apperror"
	"github.com/nitschmann/releaser/internal/helper"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

// InitService is the service interface to create a new releaser project under a path
type InitService interface {
	Call(ctx context.Context, autoYes bool) error
}

type initService struct {
	Git gitPkg.Git
}

// NewInitService initializes an new instance of InitService interface
func NewInitService(git gitPkg.Git) InitService {
	return &initService{
		Git: git,
	}
}

// Call and execute the serivce process
func (s *initService) Call(ctx context.Context, autoYes bool) error {
	commandPath := helper.CommandExecutionPathFromContext(ctx)
	// Check if git repo is initialized under current path
	err := s.checkGitRepositoryExistence(commandPath, autoYes)
	if err != nil {
		return err
	}

	// Create config Directory if not exists
	_, err = s.createConfigDir(commandPath)
	if err != nil {
		return err
	}

	return nil
}

func (s *initService) checkGitRepositoryExistence(commandPath string, autoYes bool) error {
	_, code, err := s.Git.ExecCommand(
		[]string{
			"-C",
			commandPath,
			"rev-parse",
			"--",
			"2>/dev/null",
		},
	)

	if err != nil {
		if code != 128 {
			return err
		}

		if !autoYes {
			fmt.Println(err.Error())
			msg := fmt.Sprintf("Directory '%s' is not a git respository. Do you want to continue?", commandPath)
			err = s.promptYesOrNoWithExpectedYes(msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *initService) createConfigDir(commandPath string) (string, error) {
	configDirPath := path.Join(commandPath, ".releaser")
	_, err := os.Stat(configDirPath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Creating '%s'...", configDirPath)

		return "", os.Mkdir(configDirPath, os.ModePerm)
	}

	fmt.Printf("Directory '%s' already exists - skipping", configDirPath)

	return configDirPath, nil
}

func (s *initService) promptYesOrNoWithExpectedYes(promptMsg string) error {
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
