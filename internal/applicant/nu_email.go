package applicant

import (
	"fmt"
	"strings"
)

type NUEmail string

func ParseNUEmail(str string) (NUEmail, error) {
	if ok := strings.HasSuffix(str, "@northeastern.edu"); !ok {
		return "", fmt.Errorf("invalid northeastern email. got: %s", str)
	}

	email := NUEmail(str)
	return email, nil
}

func (n *NUEmail) String() string {
	return string(*n)
}
