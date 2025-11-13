package utils

// hasPrefix checks if a string has a prefix.
// If the prefix is empty, it returns true.
// An empty string is considered to have any prefix.
func HasPrefix(s, prefix string) bool {
	if len(s) == 0 {
		return true
	}
	return s[0:len(prefix)] == prefix
}
