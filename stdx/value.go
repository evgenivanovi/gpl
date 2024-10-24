package stdx

import (
	"fmt"
	"reflect"
	"time"

	"github.com/evgenivanovi/gpl/std"
	"github.com/evgenivanovi/gpl/std/conv"
	ref "github.com/evgenivanovi/gpl/std/reflect"
	"github.com/gookit/goutil/reflects"
)

type Value interface {
	IsNil() bool
	IsEqual(Value) bool

	IsSingle() bool
	IsMultiple() bool

	IsBool() bool
	GetBool() bool
	MaybeGetBool() (bool, error)

	IsInt() bool
	GetInt() int
	MaybeGetInt() (int, error)

	IsInts() bool
	GetInts() []int
	MaybeGetInts() ([]int, error)

	IsInt64() bool
	GetInt64() int64
	MaybeGetInt64() (int64, error)

	IsInts64() bool
	GetInts64() []int64
	MaybeGetInts64() ([]int64, error)

	IsUint() bool
	GetUint() uint
	MaybeGetUint() (uint, error)

	IsUints() bool
	GetUints() []uint
	MaybeGetUints() ([]uint, error)

	IsUint64() bool
	GetUint64() uint64
	MaybeGetUint64() (uint64, error)

	IsUints64() bool
	GetUints64() []uint64
	MaybeGetUints64() ([]uint64, error)

	IsFloat64() bool
	GetFloat64() float64
	MaybeGetFloat64() (float64, error)

	IsFloats64() bool
	GetFloats64() []float64
	MaybeGetFloats64() ([]float64, error)

	IsDuration() bool
	GetDuration() time.Duration
	MaybeGetDuration() (time.Duration, error)

	IsDurations() bool
	GetDurations() []time.Duration
	MaybeGetDurations() ([]time.Duration, error)

	IsString() bool
	GetString() string
	MaybeGetString() (string, error)

	IsStrings() bool
	GetStrings() []string
	MaybeGetStrings() ([]string, error)

	// String returns the underlying value as a string representation
	String() string
}

type NilValue struct {
	err error
}

func NewNilValue(err error) NilValue {
	return NilValue{
		err: err,
	}
}

func (v NilValue) IsNil() bool {
	return true
}

func (v NilValue) IsEqual(value Value) bool {
	return false
}

func (v NilValue) IsSingle() bool {
	return false
}

func (v NilValue) IsMultiple() bool {
	return false
}

func (v NilValue) IsBool() bool {
	return false
}

func (v NilValue) GetBool() bool {
	return false
}

func (v NilValue) MaybeGetBool() (bool, error) {
	return false, v.err
}

func (v NilValue) IsInt() bool {
	return false
}

func (v NilValue) GetInt() int {
	return 0
}

func (v NilValue) MaybeGetInt() (int, error) {
	return 0, v.err
}

func (v NilValue) IsInts() bool {
	return false
}

func (v NilValue) GetInts() []int {
	return make([]int, 0)
}

func (v NilValue) MaybeGetInts() ([]int, error) {
	return make([]int, 0), v.err
}

func (v NilValue) IsInt64() bool {
	return false
}

func (v NilValue) GetInt64() int64 {
	return 0
}

func (v NilValue) MaybeGetInt64() (int64, error) {
	return 0, v.err
}

func (v NilValue) IsInts64() bool {
	return false
}

func (v NilValue) GetInts64() []int64 {
	return make([]int64, 0)
}

func (v NilValue) MaybeGetInts64() ([]int64, error) {
	return make([]int64, 0), v.err
}

func (v NilValue) IsUint() bool {
	return false
}

func (v NilValue) GetUint() uint {
	return 0
}

func (v NilValue) MaybeGetUint() (uint, error) {
	return 0, v.err
}

func (v NilValue) IsUints() bool {
	return false
}

func (v NilValue) GetUints() []uint {
	return make([]uint, 0)
}

func (v NilValue) MaybeGetUints() ([]uint, error) {
	return make([]uint, 0), v.err
}

func (v NilValue) IsUint64() bool {
	return false
}

func (v NilValue) GetUint64() uint64 {
	return 0
}

func (v NilValue) MaybeGetUint64() (uint64, error) {
	return 0, v.err
}

func (v NilValue) IsUints64() bool {
	return false
}

func (v NilValue) GetUints64() []uint64 {
	return make([]uint64, 0)
}

func (v NilValue) MaybeGetUints64() ([]uint64, error) {
	return make([]uint64, 0), v.err
}

func (v NilValue) IsFloat64() bool {
	return false
}

func (v NilValue) GetFloat64() float64 {
	return 0
}

func (v NilValue) MaybeGetFloat64() (float64, error) {
	return 0, v.err
}

func (v NilValue) IsFloats64() bool {
	return false
}

func (v NilValue) GetFloats64() []float64 {
	return make([]float64, 0)
}

func (v NilValue) MaybeGetFloats64() ([]float64, error) {
	return make([]float64, 0), v.err
}

func (v NilValue) IsDuration() bool {
	return false
}

func (v NilValue) GetDuration() time.Duration {
	return 0
}

func (v NilValue) MaybeGetDuration() (time.Duration, error) {
	return 0, v.err
}

func (v NilValue) IsDurations() bool {
	return false
}

func (v NilValue) GetDurations() []time.Duration {
	return make([]time.Duration, 0)
}

func (v NilValue) MaybeGetDurations() ([]time.Duration, error) {
	return make([]time.Duration, 0), v.err
}

func (v NilValue) IsString() bool {
	return false
}

func (v NilValue) GetString() string {
	return ""
}

func (v NilValue) MaybeGetString() (string, error) {
	return "", v.err
}

func (v NilValue) IsStrings() bool {
	return false
}

func (v NilValue) GetStrings() []string {
	return make([]string, 0)
}

func (v NilValue) MaybeGetStrings() ([]string, error) {
	return make([]string, 0), v.err
}

func (v NilValue) String() string {
	return ""
}

func (v NilValue) Error() error {
	return v.err
}

type DefaultValue struct {
	value  any
	values []any

	isBool bool

	isInt  bool
	isInts bool

	isInt64  bool
	isInts64 bool

	isUint  bool
	isUints bool

	isUint64  bool
	isUints64 bool

	isFloat64  bool
	isFloats64 bool

	isDuration  bool
	isDurations bool

	isString  bool
	isStrings bool
}

func (v DefaultValue) IsNil() bool {
	return v.isSingleValueNil() && v.isMultipleValuesNil()
}

func (v DefaultValue) IsEqual(other Value) bool {
	return v.String() == other.String()
}

func (v DefaultValue) IsSingle() bool {
	return !v.isSingleValueNil()
}

func (v DefaultValue) IsMultiple() bool {
	return !v.isMultipleValuesNil()
}

func (v DefaultValue) isSingleValueNil() bool {
	return v.value == nil
}

func (v DefaultValue) isMultipleValuesNil() bool {
	return len(v.values) == 0 || len(v.values) == 1 && v.values[0] == nil
}

func (v DefaultValue) IsBool() bool {
	return v.isBool
}

func (v DefaultValue) GetBool() bool {
	val, _ := v.MaybeGetBool()
	return val
}

func (v DefaultValue) MaybeGetBool() (bool, error) {
	raw := v.stringForSingle()
	return conv.MapBool(raw)
}

func (v DefaultValue) IsInt() bool {
	return v.isInt
}

func (v DefaultValue) GetInt() int {
	val, _ := v.MaybeGetInt()
	return val
}

func (v DefaultValue) MaybeGetInt() (int, error) {
	raw := v.stringForSingle()
	return conv.MapInt(raw)
}

func (v DefaultValue) IsInts() bool {
	return v.isInts
}

func (v DefaultValue) GetInts() []int {
	val, _ := v.MaybeGetInts()
	return val
}

func (v DefaultValue) MaybeGetInts() ([]int, error) {
	if !v.IsInts() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with int type",
		)
	}

	result := make([]int, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(int); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with int type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsInt64() bool {
	return v.isInt64
}

func (v DefaultValue) GetInt64() int64 {
	val, _ := v.MaybeGetInt64()
	return val
}

func (v DefaultValue) MaybeGetInt64() (int64, error) {
	raw := v.stringForSingle()
	return conv.MapInt64(raw)
}

func (v DefaultValue) IsInts64() bool {
	return v.isInts64
}

func (v DefaultValue) GetInts64() []int64 {
	val, _ := v.MaybeGetInts64()
	return val
}

func (v DefaultValue) MaybeGetInts64() ([]int64, error) {
	if !v.IsInts64() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with int64 type",
		)
	}

	result := make([]int64, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(int64); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with int64 type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsUint() bool {
	return v.isUint
}

func (v DefaultValue) GetUint() uint {
	val, _ := v.MaybeGetUint()
	return val
}

func (v DefaultValue) MaybeGetUint() (uint, error) {
	raw := v.stringForSingle()
	return conv.MapUint(raw)
}

func (v DefaultValue) IsUints() bool {
	return v.isUints
}

func (v DefaultValue) GetUints() []uint {
	val, _ := v.MaybeGetUints()
	return val
}

func (v DefaultValue) MaybeGetUints() ([]uint, error) {
	if !v.IsUints() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with uint type",
		)
	}

	result := make([]uint, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(uint); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with uint type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsUint64() bool {
	return v.isUint64
}

func (v DefaultValue) GetUint64() uint64 {
	val, _ := v.MaybeGetUint64()
	return val
}

func (v DefaultValue) MaybeGetUint64() (uint64, error) {
	raw := v.stringForSingle()
	return conv.MapUint64(raw)
}

func (v DefaultValue) IsUints64() bool {
	return v.isUints64
}

func (v DefaultValue) GetUints64() []uint64 {
	val, _ := v.MaybeGetUints64()
	return val
}

func (v DefaultValue) MaybeGetUints64() ([]uint64, error) {
	if !v.IsUints64() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with uint64 type",
		)
	}

	result := make([]uint64, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(uint64); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with uint64 type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsFloat64() bool {
	return v.isFloat64
}

func (v DefaultValue) GetFloat64() float64 {
	val, _ := v.MaybeGetFloat64()
	return val
}

func (v DefaultValue) MaybeGetFloat64() (float64, error) {
	raw := v.stringForSingle()
	return conv.MapFloat64(raw)
}

func (v DefaultValue) IsFloats64() bool {
	return v.isFloats64
}

func (v DefaultValue) GetFloats64() []float64 {
	val, _ := v.MaybeGetFloats64()
	return val
}

func (v DefaultValue) MaybeGetFloats64() ([]float64, error) {
	if !v.IsFloats64() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with float64 type",
		)
	}

	result := make([]float64, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(float64); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with float64 type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsDuration() bool {
	return v.isDuration
}

func (v DefaultValue) GetDuration() time.Duration {
	val, _ := v.MaybeGetDuration()
	return val
}

func (v DefaultValue) MaybeGetDuration() (time.Duration, error) {
	raw := v.stringForSingle()
	return conv.MapDuration(raw)
}

func (v DefaultValue) IsDurations() bool {
	return v.isDurations
}

func (v DefaultValue) GetDurations() []time.Duration {
	val, _ := v.MaybeGetDurations()
	return val
}

func (v DefaultValue) MaybeGetDurations() ([]time.Duration, error) {
	if !v.IsDurations() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with time.Duration type",
		)
	}

	result := make([]time.Duration, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(time.Duration); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with time.Duration type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) IsString() bool {
	return v.isString
}

func (v DefaultValue) GetString() string {
	val, _ := v.MaybeGetString()
	return val
}

func (v DefaultValue) MaybeGetString() (string, error) {
	return v.stringForSingle(), nil
}

func (v DefaultValue) IsStrings() bool {
	return v.isStrings
}

func (v DefaultValue) GetStrings() []string {
	val, _ := v.MaybeGetStrings()
	return val
}

func (v DefaultValue) MaybeGetStrings() ([]string, error) {
	if !v.IsStrings() {
		return nil, fmt.Errorf(
			"values cannot be parsed as slice with string type",
		)
	}

	result := make([]string, len(v.values))

	for i, value := range v.values {
		if cast, ok := value.(string); ok {
			result[i] = cast
		} else {
			return nil, fmt.Errorf(
				"values cannot be parsed as slice with string type",
			)
		}
	}

	return result, nil
}

func (v DefaultValue) String() string {
	if v.IsNil() {
		return ""
	}

	if v.IsSingle() {
		return v.stringForSingle()
	}

	if v.IsMultiple() {
		return v.stringForMultiple()
	}

	return ""
}

func (v DefaultValue) stringForSingle() string {

	if v.IsNil() {
		return ""
	}

	if v.IsSingle() {
		return fmt.Sprintf("%v", v.value)
	}

	return ""

}

func (v DefaultValue) stringForMultiple() string {
	if v.IsNil() {
		return ""
	}

	if v.IsMultiple() {
		return fmt.Sprintf("%v", v.values)
	}

	return ""
}

// NewValue creates a new instance of NewValue
func NewValue(value interface{}) Value {

	if value == nil {
		return DefaultValue{
			value:  nil,
			values: nil,
		}
	}

	rv := reflect.ValueOf(value)
	rvt := rv.Type()
	rvk := rv.Kind()

	if !reflects.IsArrayOrSlice(rvk) {

		if rvk == reflect.Bool {
			return DefaultValue{
				isBool: true,
				value:  value,
				values: nil,
			}
		}

		if rvk == reflect.Int {
			return DefaultValue{
				isInt:  true,
				value:  value,
				values: nil,
			}
		}

		if rvk == reflect.Int64 {
			return DefaultValue{
				isInt64: true,
				value:   value,
				values:  nil,
			}
		}

		if rvk == reflect.Uint {
			return DefaultValue{
				isUint: true,
				value:  value,
				values: nil,
			}
		}

		if rvk == reflect.Uint64 {
			return DefaultValue{
				isUint64: true,
				value:    value,
				values:   nil,
			}
		}

		if rvk == reflect.Float64 {
			return DefaultValue{
				isFloat64: true,
				value:     value,
				values:    nil,
			}
		}

		if ref.IsDurationType(rvt) {
			return DefaultValue{
				isDuration: true,
				value:      value,
				values:     nil,
			}
		}

		if rvk == reflect.String {
			return DefaultValue{
				isString: true,
				value:    value,
				values:   nil,
			}
		}

	}

	if reflects.IsArrayOrSlice(rvk) {

		switch actual := value.(type) {
		case []int:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isInts: true,
					value:  nil,
					values: origin,
				}
			}
		case []int64:
			origin := make([]interface{}, len(actual))
			for ind, val := range actual {
				origin[ind] = val
			}
			return DefaultValue{
				isInts64: true,
				value:    nil,
				values:   origin,
			}
		case []uint:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isUints: true,
					value:   nil,
					values:  origin,
				}
			}
		case []uint64:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isUints64: true,
					value:     nil,
					values:    origin,
				}
			}
		case []float64:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isFloats64: true,
					value:      nil,
					values:     origin,
				}
			}
		case []time.Duration:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isDurations: true,
					value:       nil,
					values:      origin,
				}
			}
		case []string:
			{
				origin := make([]interface{}, len(actual))
				for ind, val := range actual {
					origin[ind] = val
				}
				return DefaultValue{
					isStrings: true,
					value:     nil,
					values:    origin,
				}
			}
		}

	}

	return DefaultValue{
		value:  nil,
		values: nil,
	}

}

func ValueTypeName(value Value) string {
	if def, ok := value.(DefaultValue); ok {
		if def.value != nil {
			return reflect.TypeOf(def.value).Name()
		} else if def.values != nil {
			return reflect.TypeOf(def.values[0]).Name()
		}
	}

	if _, ok := value.(NilValue); ok {
		return std.Nil
	}

	return std.Empty
}
