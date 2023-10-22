package std

func If[T any](cond bool, onTrue, onFalse T) T {
	if cond {
		return onTrue
	}
	return onFalse
}

func IfInvoke[T any](cond bool, onTrue func() T, onFalse func() T) T {
	if cond {
		return onTrue()
	}
	return onFalse()
}
