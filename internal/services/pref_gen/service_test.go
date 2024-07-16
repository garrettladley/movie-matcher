package pref_gen_test

import (
	"math/rand"
	"testing"

	"movie-matcher/internal/services/pref_gen"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePeople(t *testing.T) {
	t.Parallel()
	t.Run("returns the correct number of people", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			count uint
		}{
			{count: 0},
			{count: 1},
			{count: 8},
		}
		for _, tc := range testCases {
			assert.Len(t, pref_gen.GeneratePeople(rand.New(rand.NewSource(0)), tc.count), int(tc.count))
		}
	})
}
