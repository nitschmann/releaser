package config

// Release has release specific config settings
type Release struct {
	FirstTag          *string `mapstructure:"first_tag" yaml:"first_tag" validate:"required"`
	Target            *string `mapstructure:"target" yaml:"target" validate:"required"`
	NameFormat        *string `mapstructure:"name_format" yaml:"name_format" validate:"required"`
	DescriptionFormat *string `mapstructure:"description_format" yaml:"description_format" validate:"required"`
}

func newRelease() Release {
	return Release{
		DescriptionFormat: &ReleaseDescriptionFormatDefault,
		FirstTag:          &ReleaseFirstTagDefault,
		NameFormat:        &ReleaseNameFormatDefault,
		Target:            &ReleaseTargetDefault,
	}
}

// GetDescriptionFormat returns the value of the DescriptionFormat field if given, else default value
func (r Release) GetDescriptionFormat() string {
	if r.DescriptionFormat != nil {
		return *r.DescriptionFormat
	}

	return ReleaseDescriptionFormatDefault
}

// GetFirstTag returns the value of the FirstTag field if given, else default value
func (r Release) GetFirstTag() string {
	if r.FirstTag != nil {
		return *r.FirstTag
	}

	return ReleaseFirstTagDefault
}

// GetNameFormat returns the value of the NameFormat field if given, else default value
func (r Release) GetNameFormat() string {
	if r.NameFormat != nil {
		return *r.NameFormat
	}

	return ReleaseNameFormatDefault
}

// GetTarget returns the value of the Target field if given, else default value
func (r Release) GetTarget() string {
	if r.Target != nil {
		return *r.Target
	}

	return ReleaseTargetDefault
}
