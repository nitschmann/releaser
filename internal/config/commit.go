package config

// Commit has git commit specific config settings
type Commit struct {
	AllowedWithoutType *bool    `mapstructure:"allowed_without_type" yaml:"allowed_without_type" validate:"required"`
	MessageFormat      *string  `mapstructure:"message_format" yaml:"message_format" validate:"required"`
	Types              []string `mapstructure:"types" yaml:"types" validate:"required,unique,dive,alphanum"`
}

func newCommit() Commit {
	return Commit{
		AllowedWithoutType: &CommitAllowedWithoutTypeDefault,
		MessageFormat:      &CommitMessageFormatDefault,
		Types:              CommitTypesDefault,
	}
}

// GetAllowedWithoutType returns the value of the AllowedWithoutType field if given, else default value
func (c Commit) GetAllowedWithoutType() bool {
	if c.AllowedWithoutType != nil {
		return *c.AllowedWithoutType
	}

	return CommitAllowedWithoutTypeDefault
}

// GetMessageFormat returns the value of the MessageFormat field  if given, else default value
func (c Commit) GetMessageFormat() string {
	if c.MessageFormat != nil {
		return *c.MessageFormat
	}

	return CommitMessageFormatDefault
}

// Types  the value of the Types field  if given, else default value
func (c Commit) GetTypes() []string {
	if len(c.Types) > 0 {
		return c.Types
	}

	return CommitTypesDefault
}
