package translate

import (
	"regexp"
	"strings"
)

// Empty is used to check if string is empty
func Empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// NotEmpty check if string is not empty
func NotEmpty(s string) bool {
	return !Empty(s)
}

//IsEmail check if string is valid email
func IsEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}
