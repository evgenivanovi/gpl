package reflect

import (
	"reflect"
	"time"

	"github.com/evgenivanovi/gpl/std/ptr"
)

var durationType = reflect.TypeOf(time.Duration(0))
var durationPointerType = reflect.TypeOf(ptr.Duration(time.Duration(0)))

var timeType = reflect.TypeOf(time.Time{})
var timePointerType = reflect.TypeOf(&time.Time{})

func IsDuration(value any) bool {
	_, ok := value.(time.Duration)
	return ok
}

func IsDurationValue(value reflect.Value) bool {
	typ := value.Type()
	return IsDurationType(typ)
}

func IsDurationType(value reflect.Type) bool {
	return value == durationType
}

func IsDurationPointerType(value reflect.Type) bool {
	return value == durationPointerType
}

func IsTime(value any) bool {
	_, ok := value.(time.Time)
	return ok
}

func IsTimeValue(value reflect.Value) bool {
	typ := value.Type()
	return IsTimeType(typ) || IsTimePointerType(typ)
}

func IsTimeType(value reflect.Type) bool {
	return value == timeType
}

func IsTimePointerType(value reflect.Type) bool {
	return value == timePointerType
}
