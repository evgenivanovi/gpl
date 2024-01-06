package cfg

import (
	"time"

	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx"
	"github.com/gookit/goutil/strutil"
)

/* __________________________________________________ */

func FirstBoolOr(or bool) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			sourceValue := source.Map(
				func(raw string) (any, error) {
					return conv.MapBool(raw)
				},
			)
			if sourceValue.IsBool() {
				return sourceValue
			}
		}
		return result
	}
}

/* __________________________________________________ */

func FirstInt(or int) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			sourceValue := source.Map(
				func(raw string) (any, error) {
					return conv.MapInt(raw)
				},
			)
			if sourceValue.IsInt() {
				return sourceValue
			}
		}
		return result
	}
}

func FirstInt64(or int64) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			sourceValue := source.Map(
				func(raw string) (any, error) {
					return conv.MapInt64(raw)
				},
			)
			if sourceValue.IsInt64() {
				return sourceValue
			}
		}
		return result
	}
}

/* __________________________________________________ */

func FirstDurationOr(or time.Duration) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			sourceValue := source.Map(
				func(raw string) (any, error) {
					return conv.MapDuration(raw)
				},
			)
			if sourceValue.IsDuration() {
				return sourceValue
			}
		}
		return result
	}
}

/* __________________________________________________ */

func FirstStringOr(or string) func(sources []Source) string {
	return func(sources []Source) string {
		result := or
		for _, source := range sources {
			sourceResult, sourcePresent := source.Get()
			if sourcePresent {
				return sourceResult
			}
		}
		return result
	}
}

func FirstStringNotEmpty(or string) func(sources []Source) string {
	return func(sources []Source) string {
		result := or
		for _, source := range sources {
			sourceResult, sourcePresent := source.Get()
			if sourcePresent && strutil.IsNotBlank(sourceResult) {
				return sourceResult
			}
		}
		return result
	}
}

func FirstStringNotEmptyElse() func(sources []Source) (string, error) {
	return func(sources []Source) (string, error) {
		var result string
		for _, source := range sources {
			sourceResult, sourcePresent := source.Get()
			if sourcePresent && strutil.IsNotBlank(sourceResult) {
				return sourceResult, nil
			}
		}
		return result, PropertyNotFoundError
	}
}

func FirstStringNotEmptyThrow() func(sources []Source) (string, error) {
	return func(sources []Source) (string, error) {
		for _, source := range sources {
			sourceResult, sourcePresent := source.Get()
			if sourcePresent && strutil.IsNotBlank(sourceResult) {
				return sourceResult, nil
			}
		}
		panic(PropertyNotFoundError)
	}
}

/* __________________________________________________ */
