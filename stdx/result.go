package stdx

type Result[T any] struct {
	Value T
	Err   error
}

func NewResult[T any](action func() (T, error)) Result[T] {
	value, err := action()

	if err != nil {
		return NewResultError[T](err)
	}

	return NewResultValue[T](value)
}

func NewResultValue[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}

func NewResultError[T any](err error) Result[T] {
	return Result[T]{
		Err: err,
	}
}

func (r *Result[T]) IsSuccess() bool {
	return r.Err == nil
}

func (r *Result[T]) IsFailure() bool {
	return r.Err != nil
}

func (r *Result[T]) MustValue() T {
	if r.IsSuccess() {
		return r.Value
	}
	panic(r.Err)
}

func (r *Result[T]) MustError() error {
	if r.IsFailure() {
		return r.Err
	}
	return nil
}

func MapResult[T, R any](result *Result[T], mapping func(T) R) *Result[R] {
	if result == nil {
		return nil
	}

	if result.IsSuccess() {
		return &Result[R]{
			Value: mapping(result.Value),
		}
	}

	return &Result[R]{
		Err: result.Err,
	}
}

func FlatMapResult[T, R any](result *Result[T], mapping func(T) *Result[R]) *Result[R] {
	if result == nil {
		return nil
	}

	if result.IsSuccess() {
		return mapping(result.Value)
	}

	return &Result[R]{
		Err: result.Err,
	}
}
