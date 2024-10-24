package std

import (
	"reflect"
	"testing"
)

func TestIsBoolean(t *testing.T) {
	tests := []struct {
		input  any
		expect bool
	}{
		{
			input:  nil,
			expect: false,
		},
		{
			input:  true,
			expect: true,
		},
	}

	for _, test := range tests {
		runIsTest(t, "IsBoolean", test.input, test.expect, IsBoolean)
	}
}

func TestCastToBoolean(t *testing.T) {
	tests := []struct {
		input       any
		expectCast  bool
		expectValue bool
	}{
		{
			input:       nil,
			expectCast:  false,
			expectValue: false,
		},
		{
			input:       true,
			expectCast:  true,
			expectValue: true,
		},
	}

	for _, test := range tests {
		runCastTest(t, "CastToBoolean", test.input, test.expectCast, test.expectValue, CastToBoolean)
	}
}

func TestIsString(t *testing.T) {
	tests := []struct {
		input  any
		expect bool
	}{
		{
			input:  nil,
			expect: false,
		},
		{
			input:  "",
			expect: true,
		},
	}

	for _, test := range tests {
		runIsTest(t, "IsString", test.input, test.expect, IsString)
	}
}

func TestCastToString(t *testing.T) {
	tests := []struct {
		input       any
		expectCast  bool
		expectValue string
	}{
		{
			input:       nil,
			expectCast:  false,
			expectValue: Empty,
		},
		{
			input:       "",
			expectCast:  true,
			expectValue: Empty,
		},
	}

	for _, test := range tests {
		runCastTest(t, "CastToString", test.input, test.expectCast, test.expectValue, CastToString)
	}
}

func TestIsInt32(t *testing.T) {
	tests := []struct {
		input  any
		expect bool
	}{
		{
			input:  nil,
			expect: false,
		},
		{
			input:  int32(0),
			expect: true,
		},
	}

	for _, test := range tests {
		runIsTest(t, "IsInt32", test.input, test.expect, IsInt32)
	}
}

func TestCastInt32(t *testing.T) {
	tests := []struct {
		input       any
		expectCast  bool
		expectValue int32
	}{
		{
			input:       nil,
			expectCast:  false,
			expectValue: MinusOneI,
		},
		{
			input:       int32(0),
			expectCast:  true,
			expectValue: 0,
		},
	}

	for _, test := range tests {
		runCastTest(t, "CastToInt32", test.input, test.expectCast, test.expectValue, CastToInt32)
	}
}

func TestIsInt64(t *testing.T) {
	tests := []struct {
		input  any
		expect bool
	}{
		{
			input:  nil,
			expect: false,
		},
		{
			input:  int64(0),
			expect: true,
		},
	}

	for _, test := range tests {
		runIsTest(t, "IsInt64", test.input, test.expect, IsInt64)
	}
}

func TestCastInt64(t *testing.T) {
	tests := []struct {
		input       any
		expectCast  bool
		expectValue int64
	}{
		{
			input:       nil,
			expectCast:  false,
			expectValue: MinusOneL,
		},
		{
			input:       int64(0),
			expectCast:  true,
			expectValue: 0,
		},
	}

	for _, test := range tests {
		runCastTest(t, "CastToInt64", test.input, test.expectCast, test.expectValue, CastToInt64)
	}
}

func runIsTest(t *testing.T, name string, input any, expect bool, testee func(any) bool) {
	t.Run(name, func(t *testing.T) {
		actual := testee(input)
		if expect != actual {
			t.Errorf("%s() = %v, expected %v", name, actual, expect)
		}
	})
}

func runCastTest[T any](t *testing.T, name string, input any, expectCast bool, expectValue T, testee func(any) (T, bool)) {
	t.Run(name, func(t *testing.T) {
		actualValue, actualCast := testee(input)
		if (expectCast != actualCast) || !reflect.DeepEqual(expectValue, actualValue) {
			t.Errorf("%s() = %v and %v, expected %v and %v", name, actualValue, actualCast, expectValue, expectCast)
		}
	})
}
