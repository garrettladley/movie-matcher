package applicant

import (
	"fmt"
	"regexp"
)

type NUEmail string

func ParseNUEmail(str string) (NUEmail, error) {
	re, err := regexp.Compile(`[a-zA-Z.]+@northeastern.edu`)
	if err != nil {
		return "", err
	}

	if isNUEmail := re.MatchString(str); !isNUEmail {
		return "", fmt.Errorf("invalid northeastern email. got: %s", str)
	}

	email := NUEmail(str)
	return email, nil
}

func (n *NUEmail) String() string {
	return string(*n)
}
