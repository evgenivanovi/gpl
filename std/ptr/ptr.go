package ptr

import (
	"time"
)

func Ref[T any](value T) *T {
	return &value
}

func Int(v int) *int {
	return &v
}

func Int8(v int8) *int8 {
	return &v
}

func Int16(v int16) *int16 {
	return &v
}

func Int32(v int32) *int32 {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Uint(v uint) *uint {
	return &v
}

func Uint8(v uint8) *uint8 {
	return &v
}

func Uint16(v uint16) *uint16 {
	return &v
}

func Uint32(v uint32) *uint32 {
	return &v
}

func Uint64(v uint64) *uint64 {
	return &v
}

func Float32(v float32) *float32 {
	return &v
}

func Float64(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func String(v string) *string {
	return &v
}

func Byte(v byte) *byte {
	return &v
}

func Rune(v rune) *rune {
	return &v
}

func Complex64(v complex64) *complex64 {
	return &v
}

func Complex128(v complex128) *complex128 {
	return &v
}

func Time(v time.Time) *time.Time {
	return &v
}

func Duration(v time.Duration) *time.Duration {
	return &v
}

func Slice[S ~[]T, T any](v S) *S {
	return &v
}

func Map[M ~map[K]V, K comparable, V any](v M) *M {
	return &v
}

func Channel[T any](v chan T) *chan T {
	return &v
}
