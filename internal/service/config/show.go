package config

import (
	"context"

	"gopkg.in/yaml.v3"

	configPkg "github.com/nitschmann/releaser/internal/config"
)

// ShowService  is the service interface to show a specified config as YAML encoded  []byte
type ShowService interface {
	Call(ctx context.Context, cfg configPkg.Config) ([]byte, error)
}

type showService struct{}

// NewShowService returns a new instance of ShowService interface with default values
func NewShowService() ShowService {
	return &showService{}
}

// Call and execute the service process
func (s *showService) Call(ctx context.Context, cfg configPkg.Config) ([]byte, error) {
	var yamlData []byte

	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return yamlData, err
	}
	return yamlData, nil
}
