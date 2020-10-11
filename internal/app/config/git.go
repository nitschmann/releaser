package config

// Git holds git specific config rules
type Git struct {
	Executable string `mapstructure:"executable"`
}
