package ordered_set

import (
	"slices"

	go_json "github.com/goccy/go-json"
)

// An OrderedSet is an ordered collection of elements that ensures no duplicates.
type OrderedSet[T comparable] struct {
	elems []T
}

// Creates a new OrderedSet from the provided elements by removing all duplicates. Order is preserved.
func New[T comparable](elems ...T) OrderedSet[T] {
	return OrderedSet[T]{elems: dedupe(elems)}
}

// Creates a slice from an Orderedordered_set. Order is preserved.
func (s OrderedSet[T]) Slice() []T {
	return append([]T{}, s.elems...)
}

func (s OrderedSet[T]) MarshalJSON() ([]byte, error) {
	return go_json.Marshal(s.Slice())
}

func (s *OrderedSet[T]) UnmarshalJSON(b []byte) error {
	var v []T
	if err := go_json.Unmarshal(b, &v); err != nil {
		return err
	}
	*s = New(v...)
	return nil
}

func (s *OrderedSet[T]) Len() uint {
	return uint(len(s.elems))
}

// Computes the distance between two OrderedSets, which is the sum of each element's distance.
// A single element's distance is the difference between its positions in each ordered_set.
// A set distance of 0 indicates that the sets are identical.
// Panics if the provided sets do not have the same lengths or the same elements.
func Distance[T comparable](s1 OrderedSet[T], s2 OrderedSet[T]) int {
	if len(s1.elems) != len(s2.elems) {
		return -1
	}

	distance := 0
	for i, e1 := range s1.elems {
		found := false
		for j, e2 := range s2.elems {
			if e1 == e2 {
				found = true
				distance += max(i-j, j-i, 0)
				break
			}
		}
		if !found {
			return -1
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
