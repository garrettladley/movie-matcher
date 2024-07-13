package set

import (
	"fmt"
	"slices"
)

// An OrderedSet is an ordered collection of elements that ensures no duplicates.
type OrderedSet[T comparable] struct {
	elems []T
}

// Creates a new OrderedSet from the provided elements by removing all duplicates. Order is preserved.
func New[T comparable](elems ...T) OrderedSet[T] {
	return OrderedSet[T]{elems: dedupe(elems)}
}

// Computes the distance between two OrderedSets, which is the sum of each element's distance.
// A single element's distance is the difference between its positions in each set.
// A set distance of 0 indicates that the sets are identical.
// Panics if the provided sets do not have the same lengths or the same elements.
func Distance[T comparable](s1 OrderedSet[T], s2 OrderedSet[T]) uint {
	if len(s1.elems) != len(s2.elems) {
		panic(fmt.Errorf("set lengths not equal: %d != %d", len(s1.elems), len(s2.elems)))
	}

	var distance uint = 0
	for i, e1 := range s1.elems {
		found := false
		for j, e2 := range s2.elems {
			if e1 == e2 {
				found = true
				distance += uint(max(i-j, j-i, 0))
				break
			}
		}
		if !found {
			panic(fmt.Errorf("set elements are not identical"))
		}
	}
	return distance
}

// Creates a new slice from the given slice with duplicate elements removed.
func dedupe[T comparable](old []T) []T {
	new := []T{}
	for _, elem := range old {
		if !slices.Contains(new, elem) {
			new = append(new, elem)
		}
	}
	return new
}
