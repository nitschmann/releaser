package config

// Branch has git branch specific config settings
type Branch struct {
	AllowedWithoutType *bool    `mapstructure:"allowed_without_type" yaml:"allowed_without_type" validate:"required"`
	Delimiter          *string  `mapstructure:"delimiter" yaml:"delimiter" validate:"required,len=1"`
	NameFormat         *string  `mapstructure:"name_format" yaml:"name_format" validate:"required`
	Types              []string `mapstructure:"types" yaml:"types" validate:"required,unique,dive,lowercase,alphanum"`
}

func newBranch() Branch {
	return Branch{
		AllowedWithoutType: &BranchAllowedWithoutTypeDefault,
		Delimiter:          &BranchDelimiterDefault,
		NameFormat:         &BranchNameFormatDefault,
		Types:              BranchTypesDefault,
	}
}

// GetAllowedWithoutType returns the value of the AllowedWithoutType field if present, else default value
func (b Branch) GetAllowedWithoutType() bool {
	if b.AllowedWithoutType != nil {
		return *b.AllowedWithoutType
	}

	return BranchAllowedWithoutTypeDefault
}

// GetDelimiter returns the value of the Delimiter field if present, else default value
func (b Branch) GetDelimiter() string {
	if b.Delimiter != nil {
		return *b.Delimiter
	}

	return BranchDelimiterDefault
}

// GetNameFormat returns the value of the NameFormat field if present, else default value
func (b Branch) GetNameFormat() string {
	if b.NameFormat != nil {
		return *b.NameFormat
	}

	return BranchNameFormatDefault
}

// GetTypes returns the value of the Types field if present, else default value
func (b Branch) GetTypes() []string {
	if len(b.Types) > 0 {
		return b.Types
	}

	return BranchTypesDefault
}
