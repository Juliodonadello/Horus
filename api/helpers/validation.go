package helpers

// ValidString takes a string and a slice of string and returns true if the string is present in the slice.
func ValidString(s string, v []string) bool {
	for i := 0; i < len(v); i++ {
		if v[i] == s {
			return true
		}
	}
	return false
}
