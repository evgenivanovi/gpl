package std

import (
	"testing"
)

/* __________________________________________________ */

type IsBooleanData struct {
	Input  any
	Expect bool
}

var IsBooleanTestData = []IsBooleanData{
	IsBooleanData{
		Input:  nil,
		Expect: false,
	},
	IsBooleanData{
		Input:  true,
		Expect: true,
	},
}

func TestIsBoolean(t *testing.T) {

	// given && when &&then
	for _, testData := range IsBooleanTestData {
		t.Run(
			"TestIsBoolean",
			func(t *testing.T) {
				actual := IsBoolean(testData.Input)
				if testData.Expect != actual {
					t.Errorf("IsBoolean() = %v, expected %v", actual, testData.Expect)
				}
			},
		)
	}

}

/* __________________________________________________ */

type CastToBooleanData struct {
	Input       any
	ExpectCast  bool
	ExpectValue bool
}

var CastToBooleanTestData = []CastToBooleanData{
	CastToBooleanData{
		Input:       nil,
		ExpectCast:  false,
		ExpectValue: false,
	},
	CastToBooleanData{
		Input:       true,
		ExpectCast:  true,
		ExpectValue: true,
	},
}

func TestCastToBoolean(t *testing.T) {

	// given && when &&then
	for _, testData := range CastToBooleanTestData {
		t.Run(
			"TestCastToBoolean",
			func(t *testing.T) {

				actualValue, actualCast := CastToBoolean(testData.Input)

				if (testData.ExpectCast != actualCast) || (testData.ExpectValue != actualValue) {
					t.Errorf(
						"CastToBoolean() = %v and %v, expected %v and %v",
						actualValue, actualValue, testData.ExpectCast, testData.ExpectValue,
					)
				}

			},
		)
	}

}

/* __________________________________________________ */

type IsStringData struct {
	Input  any
	Expect bool
}

var IsStringTestData = []IsStringData{
	IsStringData{
		Input:  nil,
		Expect: false,
	},
	IsStringData{
		Input:  "",
		Expect: true,
	},
}

func TestIsString(t *testing.T) {

	// given && when &&then
	for _, testData := range IsStringTestData {
		t.Run(
			"TestIsString",
			func(t *testing.T) {
				actual := IsString(testData.Input)
				if testData.Expect != actual {
					t.Errorf("IsString() = %v, expected %v", actual, testData.Expect)
				}
			},
		)
	}

}

/* __________________________________________________ */

type CastToStringData struct {
	Input       any
	ExpectCast  bool
	ExpectValue string
}

var CastToStringTestData = []CastToStringData{
	CastToStringData{
		Input:       nil,
		ExpectCast:  false,
		ExpectValue: Empty,
	},
	CastToStringData{
		Input:       "",
		ExpectCast:  true,
		ExpectValue: Empty,
	},
}

func TestCastToString(t *testing.T) {

	// given && when &&then
	for _, testData := range CastToStringTestData {
		t.Run(
			"TestCastToString",
			func(t *testing.T) {

				actualValue, actualCast := CastToString(testData.Input)

				if (testData.ExpectCast != actualCast) || (testData.ExpectValue != actualValue) {
					t.Errorf(
						"CastToString() = %v and %v, expected %v and %v",
						actualValue, actualValue, testData.ExpectCast, testData.ExpectValue,
					)
				}

			},
		)
	}

}

/* __________________________________________________ */

type IsInt32Data struct {
	Input  any
	Expect bool
}

var IsInt32TestData = []IsInt32Data{
	IsInt32Data{
		Input:  nil,
		Expect: false,
	},
	IsInt32Data{
		Input:  int32(0),
		Expect: true,
	},
}

func TestIsInt32(t *testing.T) {

	// given && when &&then
	for _, testData := range IsInt32TestData {
		t.Run(
			"TestIsInt32",
			func(t *testing.T) {
				actual := IsInt32(testData.Input)
				if testData.Expect != actual {
					t.Errorf("IsInt32() = %v, expected %v", actual, testData.Expect)
				}
			},
		)
	}

}

/* __________________________________________________ */

type CastToInt32Data struct {
	Input       any
	ExpectCast  bool
	ExpectValue int32
}

var CastInt32TestData = []CastToInt32Data{
	CastToInt32Data{
		Input:       nil,
		ExpectCast:  false,
		ExpectValue: MinusOneI,
	},
	CastToInt32Data{
		Input:       int32(0),
		ExpectCast:  true,
		ExpectValue: 0,
	},
}

func TestCastInt32(t *testing.T) {

	// given && when &&then
	for _, testData := range CastInt32TestData {
		t.Run(
			"TestCastInt32",
			func(t *testing.T) {

				actualValue, actualCast := CastToInt32(testData.Input)

				if (testData.ExpectCast != actualCast) || (testData.ExpectValue != actualValue) {
					t.Errorf(
						"CastInt32() = %v and %v, expected %v and %v",
						actualValue, actualValue, testData.ExpectCast, testData.ExpectValue,
					)
				}

			},
		)
	}

}

/* __________________________________________________ */

type IsInt64Data struct {
	Input  any
	Expect bool
}

var IsInt64TestData = []IsInt64Data{
	IsInt64Data{
		Input:  nil,
		Expect: false,
	},
	IsInt64Data{
		Input:  int64(0),
		Expect: true,
	},
}

func TestIsInt64(t *testing.T) {

	// given && when &&then
	for _, testData := range IsInt64TestData {
		t.Run(
			"TestIsInt64",
			func(t *testing.T) {
				actual := IsInt64(testData.Input)
				if testData.Expect != actual {
					t.Errorf("IsInt64() = %v, expected %v", actual, testData.Expect)
				}
			},
		)
	}

}

/* __________________________________________________ */

type CastToInt64Data struct {
	Input       any
	ExpectCast  bool
	ExpectValue int64
}

var CastInt64TestData = []CastToInt64Data{
	CastToInt64Data{
		Input:       nil,
		ExpectCast:  false,
		ExpectValue: MinusOneL,
	},
	CastToInt64Data{
		Input:       int64(0),
		ExpectCast:  true,
		ExpectValue: 0,
	},
}

func TestCastInt64(t *testing.T) {

	// given && when &&then
	for _, testData := range CastInt64TestData {
		t.Run(
			"TestCastInt64",
			func(t *testing.T) {

				actualValue, actualCast := CastToInt64(testData.Input)

				if (testData.ExpectCast != actualCast) || (testData.ExpectValue != actualValue) {
					t.Errorf(
						"CastInt64() = %v and %v, expected %v and %v",
						actualValue, actualValue, testData.ExpectCast, testData.ExpectValue,
					)
				}

			},
		)
	}

}

/* __________________________________________________ */
