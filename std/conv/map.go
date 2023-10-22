package conv

import (
	"fmt"
	"strconv"
	"time"

	"github.com/evgenivanovi/gpl/std"
)

type FromBool func(bool) string
type ToBool func(string) (bool, error)

func BoolString(value bool) string {
	return strconv.FormatBool(value)
}

func MapBool(raw string) (bool, error) {
	val, err := strconv.ParseBool(raw)
	if err != nil {
		return false, fmt.Errorf("value cannot be parsed as bool: %v", err)
	}
	return val, nil
}

func MustMapBool(raw string) bool {
	return std.Must(MapBool(raw))
}

func MapBoolFunc() func(raw string) (bool, error) {
	return func(raw string) (bool, error) {
		return MapBool(raw)
	}
}

type FromInt func(int) string
type ToInt func(string) (int, error)

func IntString(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func MapInt(raw string) (int, error) {
	val, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int: %v", err)
	}
	return val, nil
}

func MustMapInt(raw string) int {
	return std.Must(MapInt(raw))
}

func MapIntFunc() func(raw string) (int, error) {
	return func(raw string) (int, error) {
		return MapInt(raw)
	}
}

type FromInt64 func(int64) string
type ToInt64 func(string) (int64, error)

func Int64String(value int64) string {
	return strconv.FormatInt(value, 10)
}

func MapInt64(raw string) (int64, error) {
	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int64: %v", err)
	}
	return val, nil
}

func MustMapInt64(raw string) int64 {
	return std.Must(MapInt64(raw))
}

func MapInt64Func() func(raw string) (int64, error) {
	return func(raw string) (int64, error) {
		return MapInt64(raw)
	}
}

type FromUint func(uint) string
type ToUint func(string) (uint, error)

func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func MapUint(raw string) (uint, error) {
	val, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint: %v", err)
	}
	return uint(val), nil
}

func MustMapUint(raw string) uint {
	return std.Must(MapUint(raw))
}

func MapUintFunc() func(raw string) (uint, error) {
	return func(raw string) (uint, error) {
		return MapUint(raw)
	}
}

type FromUint16 func(uint16) string
type ToUint16 func(string) (uint16, error)

func Uint16ToString(value uint16) string {
	return strconv.FormatUint(uint64(value), 10)
}

func MapUint16(raw string) (uint16, error) {
	val, err := strconv.ParseUint(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint16: %v", err)
	}
	return uint16(val), nil
}

func MustMapUint16(raw string) uint16 {
	return std.Must(MapUint16(raw))
}

func MapUint16Func() func(raw string) (uint16, error) {
	return func(raw string) (uint16, error) {
		return MapUint16(raw)
	}
}

type FromUint64 func(uint64) string
type ToUint64 func(string) (uint64, error)

func Uint64ToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func MapUint64(raw string) (uint64, error) {
	val, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as uint64: %v", err)
	}
	return val, nil
}

func MustMapUint64(raw string) uint64 {
	return std.Must(MapUint64(raw))
}

func MapUint64Func() func(raw string) (uint64, error) {
	return func(raw string) (uint64, error) {
		return MapUint64(raw)
	}
}

type FromFloat64 func(float64) string
type ToFloat64 func(string) (float64, error)

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func MapFloat64(raw string) (float64, error) {
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as float64: %v", err)
	}
	return val, nil
}

func MustMapFloat64(raw string) float64 {
	return std.Must(MapFloat64(raw))
}

func MapFloat64Func() func(raw string) (float64, error) {
	return func(raw string) (float64, error) {
		return MapFloat64(raw)
	}
}

type FromDuration func(time.Duration) string
type ToDuration func(string) (time.Duration, error)

func DurationToString(value time.Duration) string {
	return value.String()
}

func MapDuration(raw string) (time.Duration, error) {
	val, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as time.Duration: %v", err)
	}
	return val, nil
}

func MustMapDuration(raw string) time.Duration {
	return std.Must(MapDuration(raw))
}

func MapDurationFunc() func(raw string) (time.Duration, error) {
	return func(raw string) (time.Duration, error) {
		return MapDuration(raw)
	}
}
