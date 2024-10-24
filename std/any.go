package std

import (
	"fmt"
	"reflect"
	"time"

	ref "github.com/evgenivanovi/gpl/std/reflect"
	"github.com/gookit/goutil"
)

func IsBoolean(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Bool
}

func CastToBoolean(obj interface{}) (bool, bool) {
	if goutil.IsNil(obj) {
		return false, false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Bool {

		if val, ok := obj.(*bool); ok {
			return *val, true
		}

		if val, ok := obj.(bool); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to bool"))

	}

	return false, false
}

func IsString(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.String
}

func CastToString(obj interface{}) (string, bool) {
	if goutil.IsNil(obj) {
		return Empty, false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.String {

		if val, ok := obj.(*string); ok {
			return *val, true
		}

		if val, ok := obj.(string); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to string"))

	}

	return Empty, false
}

func IsInt32(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Int32
}

func CastToInt32(obj interface{}) (int32, bool) {
	if goutil.IsNil(obj) {
		return MinusOneI, false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Int32 {

		if val, ok := obj.(*int32); ok {
			return *val, true
		}

		if val, ok := obj.(int32); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to int32"))

	}

	return MinusOneI, false
}

func IsInt64(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Int64
}

func CastToInt64(obj interface{}) (int64, bool) {
	if goutil.IsNil(obj) {
		return MinusOneL, false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Int64 {

		if val, ok := obj.(*int64); ok {
			return *val, true
		}

		if val, ok := obj.(int64); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to int64"))

	}

	return MinusOneL, false
}

func IsUint64(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Uint64
}

func CastToUint64(obj interface{}) (uint64, bool) {
	if goutil.IsNil(obj) {
		return uint64(Zero), false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Uint64 {

		if val, ok := obj.(*uint64); ok {
			return *val, true
		}

		if val, ok := obj.(uint64); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to uint64"))

	}

	return uint64(Zero), false
}

func IsFloat64(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Kind() == reflect.Float64
}

func CastToFloat64(obj interface{}) (float64, bool) {
	if goutil.IsNil(obj) {
		return float64(MinusOne), false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Float64 {

		if val, ok := obj.(*float64); ok {
			return *val, true
		}

		if val, ok := obj.(float64); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to float64"))

	}

	return float64(MinusOne), false
}

func IsDuration(obj interface{}) bool {
	if goutil.IsNil(obj) {
		return false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return ref.IsDurationType(typ)
}

func CastToDuration(obj interface{}) (time.Duration, bool) {
	if goutil.IsNil(obj) {
		return time.Duration(Zero), false
	}

	typ := reflect.TypeOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if ref.IsDurationType(typ) {

		if val, ok := obj.(*time.Duration); ok {
			return *val, true
		}

		if val, ok := obj.(time.Duration); ok {
			return val, true
		}

		panic(fmt.Errorf("unable to cast object to time.Duration"))

	}

	return time.Duration(Zero), false
}
