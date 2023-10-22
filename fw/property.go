package fw

import (
	"flag"
	"os"
	"time"

	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx"
	"github.com/gookit/goutil/strutil"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

var PropertyNotFoundError = errors.New("property not found in sources")

/* __________________________________________________ */

type Property struct {
	name    string
	sources []PropertySource
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Calc(
	calc func(sources []PropertySource) string,
) string {
	return calc(p.sources)
}

func (p *Property) CalcValue(
	calc func(sources []PropertySource) stdx.Value,
) stdx.Value {
	return calc(p.sources)
}

func (p *Property) CalcElse(
	calc func(sources []PropertySource) (string, error),
) (string, error) {
	return calc(p.sources)
}

func (p *Property) BindOne(
	source PropertySource,
) {
	p.sources = append(p.sources, source)
}

func (p *Property) BindAll(
	sources ...PropertySource,
) {
	for _, source := range sources {
		p.BindOne(source)
	}
}

func NewProperty(
	name string,
	sources ...PropertySource,
) Property {
	return Property{
		name:    name,
		sources: sources,
	}
}

/* __________________________________________________ */

type PropertySource interface {
	Get() (string, bool)
	Map(mapping func(string) (any, error)) stdx.Value
}

/* __________________________________________________ */

type ValueSource struct {
	value string
}

func (s ValueSource) Get() (string, bool) {
	return s.value, true
}

func (s ValueSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewValueSource(value string) *ValueSource {
	return &ValueSource{
		value: value,
	}
}

/* __________________________________________________ */

type ArgSource struct {
	name string
}

func (s ArgSource) Get() (string, bool) {

	if !flag.Parsed() {
		flag.Parse()
	}

	lookup := flag.Lookup(s.name)

	if lookup == nil {
		return "", false
	}

	value := lookup.Value.String()
	if value == "" {
		return "", false
	}

	return value, true

}

func (s ArgSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewArgSource(name string) *ArgSource {
	flag.String(name, "", "")
	return &ArgSource{
		name: name,
	}
}

/* __________________________________________________ */

type EnvSource struct {
	name string
}

func (s EnvSource) Get() (string, bool) {
	return os.LookupEnv(s.name)
}

func (s EnvSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewEnvSource(name string) *EnvSource {
	return &EnvSource{
		name: name,
	}
}

/* __________________________________________________ */

func FirstBoolOr(or bool) func(sources []PropertySource) stdx.Value {
	return func(sources []PropertySource) stdx.Value {
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

func FirstDurationOr(or time.Duration) func(sources []PropertySource) stdx.Value {
	return func(sources []PropertySource) stdx.Value {
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

func FirstStringOr(or string) func(sources []PropertySource) string {
	return func(sources []PropertySource) string {
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

func FirstStringNotEmpty(or string) func(sources []PropertySource) string {
	return func(sources []PropertySource) string {
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

func FirstStringNotEmptyElse() func(sources []PropertySource) (string, error) {
	return func(sources []PropertySource) (string, error) {
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

func FirstStringNotEmptyThrow() func(sources []PropertySource) (string, error) {
	return func(sources []PropertySource) (string, error) {
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
