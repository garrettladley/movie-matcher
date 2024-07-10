package model

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"
)

type ApplicantName string

var forbiddenCharacters = []rune{'/', '(', ')', '"', '<', '>', '\\', '{', '}'}

func ParseApplicantName(str string) (*ApplicantName, error) {
	if strings.TrimSpace(str) == "" {
		return nil, fmt.Errorf("name cannot be empty. got: %s", str)
	}

	if utf8.RuneCountInString(str) > 256 {
		return nil, fmt.Errorf("name is too long. got string of length: %d", len(str))
	}

	containsForbiddenCharacters := false
	for _, char := range str {
		if slices.Contains(forbiddenCharacters, char) {
			containsForbiddenCharacters = true
			break
		}
	}

	if containsForbiddenCharacters {
		return nil, fmt.Errorf("name contains forbidden characters. got: %s", str)
	}

	applicantName := ApplicantName(str)
	return &applicantName, nil
}

func (name *ApplicantName) String() string {
	return string(*name)
}
