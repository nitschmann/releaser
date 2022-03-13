package config

import "github.com/nitschmann/releaser/pkg/release/upstream"

// ReleaseUpstream has release upstream specific config settings
type ReleaseUpstream struct {
	APITokenEnvVar *string `mapstructure:"api_token_env_var" yaml:"api_token_env_var"`
}

func newReleaseUpstream() ReleaseUpstream {
	return ReleaseUpstream{}
}

// GetAPITokenEnvVar returns the value of the APITokenEnvVar field if given, else default value
func (r ReleaseUpstream) GetAPITokenEnvVar(name string) string {
	if r.APITokenEnvVar != nil {
		return *r.APITokenEnvVar
	}

	return upstream.GetRegistry().Get(name).APITokenEnvVar()
}
