package config

import "github.com/go-playground/validator/v10"

// Release is a data struct which holds configs specific for releases
type Release struct {
	FirstVersion  string `mapstructure:"first_version" validate="required"`
	LatestVersion string `mapstructure:"latest_version"`
	NewVersion    string `mapstructure:"new_version"`
}

// NewRelease returns a new pointer instance of Release with default values
func NewRelease() *Release {
	return &Release{
		FirstVersion:  "v0.0.1",
		LatestVersion: "",
		NewVersion:    "",
	}
}

// Validate runs the validators of the Release instance
func (r Release) Validate() error {
	v := validator.New()
	err := v.Struct(r)
	if err != nil {
		return err
	}

	return nil
}
