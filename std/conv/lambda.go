package conv

/* __________________________________________________ */

func Supplier[T any](value T) func() T {
	return func() T {
		return value
	}
}

func Consumer[T any](value T) func(T) {
	return func(value T) {
		return
	}
}

/* __________________________________________________ */
