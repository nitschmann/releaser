package util

// CleanList cleans a list of strings and returns it again
func CleanList(tagList []string) []string {
	var r []string
	for _, str := range tagList {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}
