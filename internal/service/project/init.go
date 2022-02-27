package project

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/nitschmann/releaser/internal/config"
	"github.com/nitschmann/releaser/internal/helper"
	"github.com/nitschmann/releaser/internal/service"
	configService "github.com/nitschmann/releaser/internal/service/config"
	gitPkg "github.com/nitschmann/releaser/pkg/git"
)

// InitService is the service interface to create a new releaser project under a path
type InitService interface {
	Call(ctx context.Context, autoYes bool) error
}

type initService struct {
	ConfigShowService configService.ShowService
	Git               gitPkg.Git
}

// NewInitService initializes an new instance of InitService interface
func NewInitService(git gitPkg.Git) InitService {
	return &initService{
		ConfigShowService: configService.NewShowService(),
		Git:               git,
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

	// Create config directory if not exists
	configDir, err := s.createConfigDir(commandPath)
	if err != nil {
		return err
	}

	// Create config file with contents
	err = s.createConfigFile(ctx, configDir, autoYes)
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
			err = service.PromptYesOrNoWithExpectedYes(msg)
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
		fmt.Printf("Creating '%s'...\n", configDirPath)

		return configDirPath, os.Mkdir(configDirPath, os.ModePerm)
	}

	fmt.Printf("Directory '%s' already exists - skipping\n", configDirPath)

	return configDirPath, nil
}

func (s *initService) createConfigFile(ctx context.Context, configDir string, autoYes bool) error {
	var err error
	configFilepath := path.Join(configDir, "config.yaml")

	if helper.FileExists(configFilepath) && !autoYes {
		msg := fmt.Sprintf("Configuration file '%s' already exists. Do you want to overwrite it?\n", configFilepath)
		err = service.PromptYesOrNoWithExpectedYes(msg)
		if err != nil {
			return err
		}
	}

	yamlData, err := s.ConfigShowService.Call(ctx, config.New())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilepath, yamlData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Sucessfully created releaser project and configuration file under " + configFilepath)

	return err
}
