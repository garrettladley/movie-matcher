package applicant

import (
	"fmt"
)

type NUID string

func ParseNUID(str string) (NUID, error) {
	if len(str) != 9 {
		return "", fmt.Errorf("NUID must be of length 9. got: %s", str)
	}

	for _, c := range str {
		if c < '0' || c > '9' {
			return "", fmt.Errorf("NUID must be a number. got: %s", str)
		}
	}
	nuid := NUID(str)
	return nuid, nil
}

func (n *NUID) String() string {
	return string(*n)
}
