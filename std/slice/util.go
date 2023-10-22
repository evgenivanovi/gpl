package slices

func IsEmpty[E any](slice []E) bool {
	return len(slice) == 0
}

func IsNotEmpty[E any](slice []E) bool {
	return len(slice) > 0
}

func IsSingle[E any](slice []E) bool {
	return len(slice) == 1
}

func IsMultiple[E any](slice []E) bool {
	return len(slice) > 1
}
