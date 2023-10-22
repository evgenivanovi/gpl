package slices

func None[S ~[]T, T any](slice S, filter func(T) bool) bool {

	if len(slice) == 0 {
		return true
	}

	for _, value := range slice {
		if filter(value) {
			return false
		}
	}

	return true

}
