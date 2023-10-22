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

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.Bool

}

func CastToBoolean(obj interface{}) (bool, bool) {

	if goutil.IsNil(obj) {
		return false, false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.Bool {

		if value, ok := obj.(*bool); ok {
			return *value, true
		}

		if value, ok := obj.(bool); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to bool"))

	}

	return false, false

}

func IsString(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.String

}

func CastToString(obj interface{}) (string, bool) {

	if goutil.IsNil(obj) {
		return Empty, false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.String {

		if value, ok := obj.(*string); ok {
			return *value, true
		}

		if value, ok := obj.(string); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to string"))

	}

	return Empty, false

}

func IsInt32(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.Int32

}

func CastToInt32(obj interface{}) (int32, bool) {

	if goutil.IsNil(obj) {
		return MinusOneI, false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.Int32 {

		if value, ok := obj.(*int32); ok {
			return *value, true
		}

		if value, ok := obj.(int32); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to int32"))

	}

	return MinusOneI, false

}

func IsInt64(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.Int64

}

func CastToInt64(obj interface{}) (int64, bool) {

	if goutil.IsNil(obj) {
		return MinusOneL, false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.Int64 {

		if value, ok := obj.(*int64); ok {
			return *value, true
		}

		if value, ok := obj.(int64); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to int64"))

	}

	return MinusOneL, false

}

func IsUint64(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.Uint64

}

func CastToUint64(obj interface{}) (uint64, bool) {

	if goutil.IsNil(obj) {
		return uint64(Zero), false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.Uint64 {

		if value, ok := obj.(*uint64); ok {
			return *value, true
		}

		if value, ok := obj.(uint64); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to uint64"))

	}

	return uint64(Zero), false

}

func IsFloat64(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.Kind() == reflect.Float64

}

func CastToFloat64(obj interface{}) (float64, bool) {

	if goutil.IsNil(obj) {
		return float64(MinusOne), false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if objType.Kind() == reflect.Float64 {

		if value, ok := obj.(*float64); ok {
			return *value, true
		}

		if value, ok := obj.(float64); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to float64"))

	}

	return float64(MinusOne), false

}

func IsDuration(obj interface{}) bool {

	if goutil.IsNil(obj) {
		return false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return ref.IsDurationType(objType)

}

func CastToDuration(obj interface{}) (time.Duration, bool) {

	if goutil.IsNil(obj) {
		return time.Duration(Zero), false
	}

	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	if ref.IsDurationType(objType) {

		if value, ok := obj.(*time.Duration); ok {
			return *value, true
		}

		if value, ok := obj.(time.Duration); ok {
			return value, true
		}

		panic(fmt.Errorf("unable to cast object to time.Duration"))

	}

	return time.Duration(Zero), false

}
