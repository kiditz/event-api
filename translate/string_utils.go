package translate

import "strings"

// Empty is used to check if string is empty
func Empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// NotEmpty check if string is not empty
func NotEmpty(s string) bool {
	return !Empty(s)
}
