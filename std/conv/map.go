package conv

import (
	"fmt"
	"strconv"
	"time"
)

/* __________________________________________________ */

type BoolMapper func(bool) string

func BoolString(value bool) string {
	return strconv.FormatBool(value)
}

type ToBoolMapper func(string) (bool, error)

func MapBool(raw string) (bool, error) {
	val, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("value cannot be parsed as bool: %v", err)
	}
	return val, nil
}

func MapBoolFunc() func(raw string) (bool, error) {
	return func(raw string) (bool, error) {
		return MapBool(raw)
	}
}

/* __________________________________________________ */

type IntMapper func(int) string

func IntString(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

type ToIntMapper func(string) (int, error)

func MapInt(raw string) (int, error) {
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int: %v", err)
	}
	return val, nil
}

func MapIntFunc() func(raw string) (int, error) {
	return func(raw string) (int, error) {
		return MapInt(raw)
	}
}

/* __________________________________________________ */

type Int64Mapper func(int64) string

func Int64String(value int64) string {
	return strconv.FormatInt(value, 10)
}

type ToInt64Mapper func(string) (int64, error)

func MapInt64(raw string) (int64, error) {
	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int64: %v", err)
	}
	return val, nil
}

func MapInt64Func() func(raw string) (int64, error) {
	return func(raw string) (int64, error) {
		return MapInt64(raw)
	}
}

/* __________________________________________________ */

type UintMapper func(uint) string

func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

type ToUintMapper func(string) (uint, error)

func MapUint(raw string) (uint, error) {
	val, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint: %v", err)
	}
	return uint(val), nil
}

func MapUintFunc() func(raw string) (uint, error) {
	return func(raw string) (uint, error) {
		return MapUint(raw)
	}
}

/* __________________________________________________ */

type Uint16Mapper func(uint16) string

func Uint16ToString(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}

type ToUint16Mapper func(string) (uint16, error)

func MapUint16(raw string) (uint16, error) {
	val, err := strconv.ParseUint(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint16: %v", err)
	}
	return uint16(val), nil
}

func MapUint16Func() func(raw string) (uint16, error) {
	return func(raw string) (uint16, error) {
		return MapUint16(raw)
	}
}

/* __________________________________________________ */

type Uint64Mapper func(uint64) string

func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

type ToUint64Mapper func(string) (uint64, error)

func MapUint64(raw string) (uint64, error) {
	val, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint64: %v", err)
	}
	return val, nil
}

func MapUint64Func() func(raw string) (uint64, error) {
	return func(raw string) (uint64, error) {
		return MapUint64(raw)
	}
}

/* __________________________________________________ */

type Float64Mapper func(float64) string

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

type ToFloat64Mapper func(string) (float64, error)

func MapFloat64(raw string) (float64, error) {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as float64: %v", err)
	}
	return val, nil
}

func MapFloat64Func() func(raw string) (float64, error) {
	return func(raw string) (float64, error) {
		return MapFloat64(raw)
	}
}

/* __________________________________________________ */

type DurationMapper func(time.Duration) string

func DurationToString(value time.Duration) string {
	return value.String()
}

type ToDurationMapper func(string) (time.Duration, error)

func MapDuration(raw string) (time.Duration, error) {
	val, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as time.Duration: %v", err)
	}
	return val, nil
}

func MapDurationFunc() func(raw string) (time.Duration, error) {
	return func(raw string) (time.Duration, error) {
		return MapDuration(raw)
	}
}

/* __________________________________________________ */
