package ordered_set_test

import (
	"testing"

	"movie-matcher/internal/ordered_set"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	t.Parallel()
	t.Run("returns the correct distance", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			set1     ordered_set.OrderedSet[string]
			set2     ordered_set.OrderedSet[string]
			distance int
		}{
			{set1: ordered_set.New[string](), set2: ordered_set.New[string](), distance: 0},
			{set1: ordered_set.New("1", "2", "3"), set2: ordered_set.New("1", "2", "3"), distance: 0},
			{set1: ordered_set.New("1", "2", "3"), set2: ordered_set.New("2", "1", "3"), distance: 2},
			{set1: ordered_set.New("1", "2", "3"), set2: ordered_set.New("3", "2", "1"), distance: 4},
			{set1: ordered_set.New("1", "2", "3"), set2: ordered_set.New("1", "2", "3"), distance: 0},
		}
		for _, tc := range testCases {
			assert.Equal(t, tc.distance, ordered_set.Distance(tc.set1, tc.set2))
		}
	})
	t.Run("-1 with invalid sets", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			set1 ordered_set.OrderedSet[string]
			set2 ordered_set.OrderedSet[string]
		}{
			{set1: ordered_set.New[string](), set2: ordered_set.New("1")},
			{set1: ordered_set.New("a"), set2: ordered_set.New("b")},
			{set1: ordered_set.New("1", "2", "3"), set2: ordered_set.New("1", "1", "3")},
		}
		for _, tc := range testCases {
			assert.Equal(t, -1, ordered_set.Distance(tc.set1, tc.set2))
		}
	})
}
