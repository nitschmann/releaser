package service

import "github.com/nitschmann/release-log/internal/app/config"

// ConfigRule is a service struct to handle config.Rule for cmd applications
type ConfigRule struct {
	currentCmdPath string
}

// NewConfigRule returns a new pointer instance of ConfigRule with the given arguments
func NewConfigRule(currentCmdPath string) *ConfigRule {
	return &ConfigRule{currentCmdPath: currentCmdPath}
}

// CurrentRule returns the current rule with applies to the path the cmd is executed in
func (c ConfigRule) CurrentRule() (config.Rule, error) {
	rule := config.Rule{}
	rule, err := config.Get().GetMatchingPathRule(c.currentCmdPath)
	if err != nil {
		return rule, err
	}

	return rule, nil
}
