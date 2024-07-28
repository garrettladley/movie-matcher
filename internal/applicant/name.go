package applicant

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Name string

const forbiddenCharacters string = "/()'\"<>\\{}"

func ParseName(str string) (Name, error) {
	if strings.TrimSpace(str) == "" {
		return "", fmt.Errorf("name cannot be empty. got: '%s'", str)
	}

	runeCountInString := utf8.RuneCountInString(str)
	if runeCountInString < 2 || runeCountInString > 256 {
		return "", fmt.Errorf("name must be between 2 and 256 characters. got string of length: %d", runeCountInString)
	}

	if strings.Contains(str, forbiddenCharacters) {
		return "", fmt.Errorf("name contains forbidden characters. got: '%s'", str)
	}

	return Name(str), nil
}

func (name *Name) String() string {
	return string(*name)
}
