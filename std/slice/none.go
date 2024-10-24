package slice

func None[S ~[]T, T any](slice S, filter func(T) bool) bool {
	if len(slice) == 0 {
		return true
	}

	for ind := range slice {
		if filter(slice[ind]) {
			return false
		}
	}

	return true
}
