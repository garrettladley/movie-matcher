package set_test

import (
	"testing"

	"movie-matcher/internal/set"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	t.Parallel()
	t.Run("returns the correct distance", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			set1     set.OrderedSet[string]
			set2     set.OrderedSet[string]
			distance uint
		}{
			{set1: set.NewOrderedSet[string](), set2: set.NewOrderedSet[string](), distance: 0},
			{set1: set.NewOrderedSet("1", "2", "3"), set2: set.NewOrderedSet("1", "2", "3"), distance: 0},
			{set1: set.NewOrderedSet("1", "2", "3"), set2: set.NewOrderedSet("2", "1", "3"), distance: 2},
			{set1: set.NewOrderedSet("1", "2", "3"), set2: set.NewOrderedSet("3", "2", "1"), distance: 4},
			{set1: set.NewOrderedSet("1", "2", "3"), set2: set.NewOrderedSet("1", "2", "3"), distance: 0},
		}
		for _, tc := range testCases {
			assert.Equal(t, tc.distance, set.Distance(tc.set1, tc.set2))
		}
	})
	t.Run("panics with invalid sets", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			set1 set.OrderedSet[string]
			set2 set.OrderedSet[string]
		}{
			{set1: set.NewOrderedSet[string](), set2: set.NewOrderedSet("1")},
			{set1: set.NewOrderedSet("a"), set2: set.NewOrderedSet("b")},
			{set1: set.NewOrderedSet("1", "2", "3"), set2: set.NewOrderedSet("1", "1", "3")},
		}
		for _, tc := range testCases {
			assert.Panics(t, func() { set.Distance(tc.set1, tc.set2) })
		}
	})
}
