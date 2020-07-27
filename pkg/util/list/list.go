package list

// CleanEmptyStrings removes empty strings from arrays or slices
func CleanEmptyStrings(list []string) []string {
	resultList := []string{}
	for _, str := range list {
		if str != "" {
			resultList = append(resultList, str)
		}
	}

	return resultList
}
