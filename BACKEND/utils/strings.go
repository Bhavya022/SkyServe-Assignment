package utils

import "strings"

// ConcatenateStrings concatenates multiple strings into a single string.
func ConcatenateStrings(strs ...string) string {
	return strings.Join(strs, "")
}

// TruncateString truncates a string to the specified length.
func TruncateString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}
