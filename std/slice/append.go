package slices

func AppendIf[S ~[]T, T any](slice S, value T, filter bool) S {
	if filter {
		return append(slice, value)
	}
	return slice
}

func AppendIfNot[S ~[]T, T any](slice S, value T, filter bool) S {
	if !filter {
		return append(slice, value)
	}
	return slice
}
