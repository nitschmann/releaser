package config

// Release is a data struct which holds configs specific for releases
type Release struct {
	FirstVersion  string `mapstructure:"first_version" validate="required"`
	LatestVersion string `mapstructure:"latest_version"`
	NewVersion    string `mapstructure:"new_version"`
}
