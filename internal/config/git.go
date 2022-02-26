package config

// Git has git specific config settings
type Git struct {
	Executable *string `mapstructure:"executable" yaml:"executable" validate:"required,alpha,lowercase"`
	Remote     *string `mapstructure:"remote" yaml:"remote" validate:"required,alphanum,lowercase"`
}

func newGit() Git {
	return Git{
		Executable: &GitExecutableDefault,
		Remote:     &GitRemoteDefault,
	}
}

// GetExecutable returns the value of the Executable field if given, else default value
func (g Git) GetExecutable() string {
	if g.Executable != nil {
		return *g.Executable
	}

	return GitExecutableDefault
}

// GetRemote returns the value of the Remote field if given, else default value
func (g Git) GetRemote() string {
	if g.Remote != nil {
		return *g.Remote
	}

	return GitRemoteDefault
}
