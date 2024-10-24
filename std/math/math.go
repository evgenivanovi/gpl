package math

import "math"

func IncInt(value int) int {
	if value == math.MaxInt {
		panic("int overflow")
	}
	return value + 1
}

func IncInt32(value int32) int32 {
	if value == math.MaxInt32 {
		panic("int32 overflow")
	}
	return value + 1
}

func IncInt64(value int64) int64 {
	if value == math.MaxInt64 {
		panic("int64 overflow")
	}
	return value + 1
}

func IncUint(value uint) uint {
	if value == math.MaxUint {
		panic("uint overflow")
	}
	return value + 1
}

func IncUint32(value uint32) uint32 {
	if value == math.MaxUint32 {
		panic("uint32 overflow")
	}
	return value + 1
}

func IncUint64(value uint64) uint64 {
	if value == math.MaxUint64 {
		panic("unit64 overflow")
	}
	return value + 1
}

func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("int overflow")
	}
	return a + b
}

func AddInt32(a, b int32) int32 {
	if a > math.MaxInt32-b {
		panic("int32 overflow")
	}
	return a + b
}

func AddInt64(a, b int64) int64 {
	if a > math.MaxInt64-b {
		panic("int64 overflow")
	}
	return a + b
}

func AddUint(a, b uint) uint {
	if a > math.MaxUint-b {
		panic("uint overflow")
	}
	return a + b
}

func AddUint32(a, b uint32) uint32 {
	if a > math.MaxUint32-b {
		panic("uint32 overflow")
	}
	return a + b
}

func AddUint64(a, b uint64) uint64 {
	if a > math.MaxUint64-b {
		panic("uint64 overflow")
	}
	return a + b
}

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt || b == math.MinInt {
		panic("int overflow")
	}
	if result/b != a {
		panic("int overflow")
	}
	return result
}

func MultiplyInt32(a, b int32) int32 {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt32 || b == math.MinInt32 {
		panic("int32 overflow")
	}
	if result/b != a {
		panic("int32 overflow")
	}
	return result
}

func MultiplyInt64(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt64 || b == math.MinInt64 {
		panic("int64 overflow")
	}
	if result/b != a {
		panic("int64 overflow")
	}
	return result
}

func Pow[T int | float64](x, y T) T {
	res := math.Pow(float64(x), float64(y))
	return T(res)
}
