package util

import (
	"github.com/stoewer/go-strcase"
)

// ToSnakeSlice converts string slice to lowercase.
func ToSnakeSlice(s []string) []string {
	ls := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		ls[i] = strcase.SnakeCase(s[i])
	}
	return ls
}
