package util

import (
	"strings"
)

// ToLowerSlice converts string slice to lowercase.
func ToLowerSlice(s []string) []string {
	ls := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ls[i] = strings.ToLower(s[i])
	}
	return ls
}
