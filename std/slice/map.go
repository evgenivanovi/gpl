package slices

// Map applies given function to every value of slice
func Map[S ~[]T, T, M any](slice S, fn func(T) M) []M {

	if slice == nil {
		return []M(nil)
	}

	if len(slice) == 0 {
		return make([]M, 0)
	}

	result := make([]M, len(slice))
	for ind, value := range slice {
		result[ind] = fn(value)
	}

	return result

}

// Mutate is like Map, but it prohibits type changes and modifies original slice.
func Mutate[S ~[]T, T any](slice S, fn func(T) T) S {

	if len(slice) == 0 {
		return slice
	}

	for ind, value := range slice {
		slice[ind] = fn(value)
	}

	return slice

}
