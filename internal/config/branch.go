package config

// Branch has git branch specific config settings
type Branch struct {
	AllowedWithoutType *bool    `mapstructure:"allowed_without_type" yaml:"allowed_without_type" validate:"required"`
	Delimiter          *string  `mapstructure:"delimiter" yaml:"delimiter" validate:"required,len=1"`
	TitleFormat        *string  `mapstructure:"title_format" yaml:"title_format" validate:"required`
	Types              []string `mapstructure:"types" yaml:"types" validate:"required,unique,dive,lowercase,alphanum"`
}

func newBranch() Branch {
	return Branch{
		AllowedWithoutType: &BranchAllowedWithoutTypeDefault,
		Delimiter:          &BranchDelimiterDefault,
		TitleFormat:        &BranchTitleFormatDefault,
		Types:              BranchTypesDefault,
	}
}

// GetAllowedWithoutType returns the value of the AllowedWithoutType field if given, else default value
func (b Branch) GetAllowedWithoutType() bool {
	if b.AllowedWithoutType != nil {
		return *b.AllowedWithoutType
	}

	return BranchAllowedWithoutTypeDefault
}

// GetDelimier returns the value of the Delimiter field if given, else default value
func (b Branch) GetDelimier() string {
	if b.Delimiter != nil {
		return *b.Delimiter
	}

	return BranchDelimiterDefault
}

// GetTitleFormat returns the value of the TitleFormat field if given, else default value
func (b Branch) GetTitleFormat() string {
	if b.TitleFormat != nil {
		return *b.TitleFormat
	}

	return BranchTitleFormatDefault
}

// GetTypes returns the value of the Types field if given, else default value
func (b Branch) GetTypes() []string {
	if len(b.Types) > 0 {
		return b.Types
	}

	return BranchTypesDefault
}
