package utilities

func IntersectionCardinality[S ~[]E, E comparable](a S, b S) uint {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	elementCount := make(map[E]int)

	for _, elem := range a {
		elementCount[elem]++
	}

	commonCount := 0
	for _, elem := range b {
		if _, exists := elementCount[elem]; exists {
			commonCount++
			delete(elementCount, elem)
		}
	}

	return uint(commonCount)
}
