package util

// Clean a list of strings
func CleanList(tagList []string) []string {
	var r []string
	for _, str := range tagList {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}
