package cmp

/* __________________________________________________ */

type Equaler interface {
	Equal(other any) bool
}

func Equal(first Equaler, second Equaler) bool {
	return first.Equal(second)
}

func NotEqual(first Equaler, second Equaler) bool {
	return !Equal(first, second)
}

func EqualAll[T Equaler](first []T, second []T) bool {

	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if !first[i].Equal(second[i]) {
			return false
		}
	}

	return true

}

func NotEqualAll[T Equaler](first []T, second []T) bool {
	return !EqualAll(first, second)
}

/* __________________________________________________ */
