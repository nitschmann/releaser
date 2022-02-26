package helper

// RemoveElementFromStringSlice removes an element at the specific index
func RemoveElementFromStringSlice(slice []string, el int) []string {
	return append(slice[:el], slice[el+1:]...)
}

// StringToPointer returns a given string as pointer
func StringToPointer(str string) *string {
	return &str
}
