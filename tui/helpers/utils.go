package helpers

// UTILITIES
// ---------

// Filter filters out empty strings from a slice of strings
func Filter(s ...string) []string {
	var filtered []string
	for _, v := range s {
		if v != "" {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
