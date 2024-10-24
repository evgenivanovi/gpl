package cfg

import (
	"time"

	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx"
	"github.com/gookit/goutil/strutil"
)

func FirstBoolOr(or bool) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			value := source.Map(
				func(raw string) (any, error) {
					return conv.MapBool(raw)
				},
			)
			if value.IsBool() {
				return value
			}
		}
		return result
	}
}

func FirstInt(or int) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			value := source.Map(
				func(raw string) (any, error) {
					return conv.MapInt(raw)
				},
			)
			if value.IsInt() {
				return value
			}
		}
		return result
	}
}

func FirstInt64(or int64) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			value := source.Map(
				func(raw string) (any, error) {
					return conv.MapInt64(raw)
				},
			)
			if value.IsInt64() {
				return value
			}
		}
		return result
	}
}

func FirstDurationOr(or time.Duration) func(sources []Source) stdx.Value {
	return func(sources []Source) stdx.Value {
		result := stdx.NewValue(or)
		for _, source := range sources {
			value := source.Map(
				func(raw string) (any, error) {
					return conv.MapDuration(raw)
				},
			)
			if value.IsDuration() {
				return value
			}
		}
		return result
	}
}

func FirstStringOr(or string) func(sources []Source) string {
	return func(sources []Source) string {
		result := or
		for _, source := range sources {
			if value, present := source.Get(); present {
				return value
			}
		}
		return result
	}
}

func FirstStringNotEmpty(or string) func(sources []Source) string {
	return func(sources []Source) string {
		result := or
		for _, source := range sources {
			if value, present := source.Get(); present && strutil.IsNotBlank(value) {
				return value
			}
		}
		return result
	}
}

func FirstStringNotEmptyElse() func(sources []Source) (string, error) {
	return func(sources []Source) (string, error) {
		var result string
		for _, source := range sources {
			if value, present := source.Get(); present && strutil.IsNotBlank(value) {
				return value, nil
			}
		}
		return result, ErrPropertyNotFound
	}
}

func FirstStringNotEmptyThrow() func(sources []Source) string {
	return func(sources []Source) string {
		for _, source := range sources {
			if value, present := source.Get(); present && strutil.IsNotBlank(value) {
				return value
			}
		}
		panic(ErrPropertyNotFound)
	}
}
