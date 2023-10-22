package slices

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// Dedup removes duplicate values from slice.
// It will alter original non-empty slice, consider copy it beforehand.
func Dedup[E constraints.Ordered](slice []E) []E {

	if len(slice) < 2 {
		return slice
	}

	slices.Sort(slice)

	temp := slice[:1]
	current := slice[0]

	for ind := 1; ind < len(slice); ind++ {
		if slice[ind] != current {
			temp = append(temp, slice[ind])
			current = slice[ind]
		}
	}

	return temp

}
