package utilities

import (
	"strings"
	"unicode"
)

func Tokenize(s string) []string {
	if s == "" {
		return []string{}
	}

	return strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
