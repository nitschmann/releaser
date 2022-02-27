package helper

// BoolPointerOrBackup returns the val attribute unless nil, else an pointer to the backupValue
func BoolPointerOrBackup(val *bool, backupValue bool) *bool {
	if val != nil {
		return val
	}

	return &backupValue
}
