package cfg

import (
	"fmt"

	"github.com/evgenivanovi/gpl/stdx"
	"github.com/pkg/errors"
)

type Property struct {
	name    string
	sources []Source
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Calc(
	calc func(sources []Source) string,
) string {
	defer p.catch()
	return calc(p.sources)
}

func (p *Property) CalcFn(
	calc func(sources []Source) string,
) func() string {
	return func() string {
		return p.Calc(calc)
	}
}

func (p *Property) CalcValue(
	calc func(sources []Source) stdx.Value,
) stdx.Value {
	defer p.catch()
	return calc(p.sources)
}

func (p *Property) CalcValueFn(
	calc func(sources []Source) stdx.Value,
) func() stdx.Value {
	return func() stdx.Value {
		return p.CalcValue(calc)
	}
}

func (p *Property) CalcElse(
	calc func(sources []Source) (string, error),
) (string, error) {
	result, err := calc(p.sources)
	if err != nil && errors.Is(err, ErrPropertyNotFound) {
		msg := fmt.Sprintf("property '%s' not found in sources", p.name)
		return "", errors.New(msg)
	}
	return result, nil
}

func (p *Property) CalcElseFn(
	calc func(sources []Source) (string, error),
) func() (string, error) {
	return func() (string, error) {
		return p.CalcElse(calc)
	}
}

func (p *Property) BindOne(
	source Source,
) {
	p.sources = append(p.sources, source)
}

func (p *Property) BindAll(
	sources ...Source,
) {
	p.sources = append(p.sources, sources...)
}

func (p *Property) catch() {
	if result := recover(); result != nil {
		if err, ok := result.(error); ok {
			if errors.Is(err, ErrPropertyNotFound) {
				msg := fmt.Sprintf("property '%s' not found in sources", p.name)
				panic(msg)
			}
		}
	}
}

func NewProperty(
	name string,
	sources ...Source,
) Property {
	return Property{
		name:    name,
		sources: sources,
	}
}
