package std

func If[T any](cond bool, yes, no T) T {
	if cond {
		return yes
	}
	return no
}

func IfInvoke[T any](cond bool, yes func() T, no func() T) T {
	if cond {
		return yes()
	}
	return no()
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
