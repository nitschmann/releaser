package helper

// CleanEmptyEntriesFromStringSlice removes empty strings from specified slice
func CleanEmptyEntriesFromStringSlice(slice []string) []string {
	var result []string

	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

// RemoveElementFromStringSlice removes an element at the specific index
func RemoveElementFromStringSlice(slice []string, el int) []string {
	return append(slice[:el], slice[el+1:]...)
}

// StringPointerOrBackup returns the val attribute unless nil, else an pointer to the backupValue
func StringPointerOrBackup(val *string, backupValue string) *string {
	if val != nil {
		return val
	}

	return &backupValue
}

// StringToPointer returns a given string as pointer
func StringToPointer(str string) *string {
	return &str
}

// StringSliceWithValuesOrBackup returns the list if not empty or the backup value
func StringSliceWithValuesOrBackup(list []string, backupValue []string) []string {
	if len(list) > 0 {
		return list
	}

	return backupValue
}

// StringSliceIncludesElement checks if a given slice of String includes the specified element
func StringSliceIncludesElement(list []string, element string) bool {
	for _, i := range list {
		if i == element {
			return true
		}
	}

	return false
}
