package applicant

import (
	"fmt"
	"regexp"
	"strings"
)

type NUEmail string

func ParseNUEmail(str string) (NUEmail, error) {
	re, err := regexp.Compile(`^[a-zA-Z]+\.[a-zA-Z]+[0-9]?@northeastern\.edu$`)
	if err != nil {
		return "", err
	}

	if isNUEmail := re.MatchString(str); !isNUEmail {
		return "", fmt.Errorf("invalid northeastern email. got: %s", str)
	}

	return NUEmail(strings.ToLower(str)), nil
}

func (n *NUEmail) String() string {
	return string(*n)
}
