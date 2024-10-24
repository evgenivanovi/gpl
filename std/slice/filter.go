package slice

import (
	"slices"
)

// Filter reduces slice values using given function.
// It operates with a copy of given slice.
func Filter[S ~[]T, T any](slice S, filter func(T) bool) S {
	if len(slice) == 0 {
		return slice
	}
	return Reduce(slices.Clone(slice), filter)
}

// Reduce is like Filter, but modifies original slice.
func Reduce[S ~[]T, T any](slice S, filter func(T) bool) S {
	if len(slice) == 0 {
		return slice
	}

	var index = 0
	for _, value := range slice {
		if filter(value) {
			slice[index] = value
			index++
		}
	}

	return slice[:index]
}
