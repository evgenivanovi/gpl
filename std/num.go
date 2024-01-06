package std

/* __________________________________________________ */

const (
	MinusOne   int = -1
	MinusTwo   int = -2
	MinusThree int = -3
	MinusFour  int = -4
	MinusFive  int = -5
	MinusSix   int = -6
	MinusSeven int = -7
	MinusEight int = -8
	MinusNine  int = -9
	MinusTen   int = -10
	Zero       int = 0
	One        int = 1
	Two        int = 2
	Three      int = 3
	Four       int = 4
	Five       int = 5
	Six        int = 6
	Seven      int = 7
	Eight      int = 8
	Nine       int = 9
	Ten        int = 10

	MinusOneI   int32 = -1
	MinusTwoI   int32 = -2
	MinusThreeI int32 = -3
	MinusFourI  int32 = -4
	MinusFiveI  int32 = -5
	MinusSixI   int32 = -6
	MinusSevenI int32 = -7
	MinusEightI int32 = -8
	MinusNineI  int32 = -9
	MinusTenI   int32 = -10
	ZeroI       int32 = 0
	OneI        int32 = 1
	TwoI        int32 = 2
	ThreeI      int32 = 3
	FourI       int32 = 4
	FiveI       int32 = 5
	SixI        int32 = 6
	SevenI      int32 = 7
	EightI      int32 = 8
	NineI       int32 = 9
	TenI        int32 = 10

	MinusOneL   int64 = -1
	MinusTwoL   int64 = -2
	MinusThreeL int64 = -3
	MinusFourL  int64 = -4
	MinusFiveL  int64 = -5
	MinusSixL   int64 = -6
	MinusSevenL int64 = -7
	MinusEightL int64 = -8
	MinusNineL  int64 = -9
	MinusTenL   int64 = -10
	ZeroL       int64 = 0
	OneL        int64 = 1
	TwoL        int64 = 2
	ThreeL      int64 = 3
	FourL       int64 = 4
	FiveL       int64 = 5
	SixL        int64 = 6
	SevenL      int64 = 7
	EightL      int64 = 8
	NineL       int64 = 9
	TenL        int64 = 10
)

/* __________________________________________________ */

func IsZeroInt32(value int32) bool {
	return value == 0
}

func IsZeroInt64(value int64) bool {
	return value == 0
}

func IsZeroInt(value int) bool {
	return value == 0
}

/* __________________________________________________ */

func IsPositiveInt32(value int32) bool {
	return value > 0
}

func IsPositiveInt64(value int64) bool {
	return value > 0
}

func IsPositiveInt(value int) bool {
	return value > 0
}

/* __________________________________________________ */

func IsNegativeInt32(value int32) bool {
	return value < 0
}

func IsNegativeInt64(value int64) bool {
	return value < 0
}

func IsNegativeInt(value int) bool {
	return value < 0
}

/* __________________________________________________ */

func IsZeroFloat32(value float32) bool {
	return value == 0
}

func IsZeroFloat64(value float64) bool {
	return value == 0
}

/* __________________________________________________ */

func IsPositiveFloat32(value float32) bool {
	return value > 0
}

func IsPositiveFloat64(value float64) bool {
	return value > 0
}

/* __________________________________________________ */

func IsNegativeFloat32(value float32) bool {
	return value < 0
}

func IsNegativeFloat64(value float64) bool {
	return value < 0
}

/* __________________________________________________ */
