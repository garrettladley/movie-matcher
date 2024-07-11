package model

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type ApplicantName string

const forbiddenCharacters string = "/()'\"<>\\{}"

func ParseApplicantName(str string) (*ApplicantName, error) {
	if strings.TrimSpace(str) == "" {
		return nil, fmt.Errorf("name cannot be empty. got: '%s'", str)
	}

	runeCountInString := utf8.RuneCountInString(str)
	if runeCountInString < 2 || runeCountInString > 256 {
		return nil, fmt.Errorf("name must be between 2 and 256 characters. got string of length: %d", runeCountInString)
	}

	if strings.Contains(str, forbiddenCharacters) {
		return nil, fmt.Errorf("name contains forbidden characters. got: '%s'", str)
	}

	applicantName := ApplicantName(str)
	return &applicantName, nil
}

func (name *ApplicantName) String() string {
	return string(*name)
}
