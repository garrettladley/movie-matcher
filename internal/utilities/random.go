package utilities

import (
	"golang.org/x/exp/rand"
)

func SelectRandom[T any](s []T, n int) []T {
	length := len(s)
	if n >= length {
		return s
	}

	indices := rand.Perm(length)[:n]

	selectedValues := make([]T, n)

	for i, index := range indices {
		selectedValues[i] = s[index]
	}

	return selectedValues
}
