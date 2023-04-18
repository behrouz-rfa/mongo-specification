package utils

func SliceIndex[T any](slice []T, predicate func(e T) bool) int {
	for i, item := range slice {
		if predicate(item) {
			return i
		}
	}

	return -1
}
