package rand

import (
	"crypto/rand"
	"encoding/binary"
)

/* __________________________________________________ */

// Int8 returns a random 8-bit signed integer.
// Return 0 and an error if unable to get random data.
func Int8() (int8, error) {
	i, err := Uint8()

	if err != nil {
		return int8(0), err
	}

	return int8(i), nil
}

// Uint8 returns a random 8-bit unsigned integer.
// Return 0 and an error if unable to get rand data.
func Uint8() (uint8, error) {
	var bytes [1]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint8(0), err
	}

	return bytes[0], nil
}

/* __________________________________________________ */

// Int16 returns a random 16-bit signed integer.
// Return 0 and an error if unable to get rand data.
func Int16() (int16, error) {
	i, err := Uint16()

	if err != nil {
		return int16(0), err
	}

	return int16(i), nil
}

// Uint16 returns a random 16-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func Uint16() (uint16, error) {
	var bytes [2]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint16(0), err
	}

	return binary.LittleEndian.Uint16(bytes[:]), nil
}

/* __________________________________________________ */

// Int32 returns a random 32-bit signed integer.
// Return 0 and an error if unable to get random data.
func Int32() (int32, error) {
	i, err := Uint32()

	if err != nil {
		return int32(0), err
	}

	return int32(i), nil
}

// Uint32 returns a rand 32-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func Uint32() (uint32, error) {
	var bytes [4]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint32(0), err
	}

	return binary.LittleEndian.Uint32(bytes[:]), nil
}

/* __________________________________________________ */

// Int64 returns a random 64-bit signed integer.
// Return 0 and an error if unable to get random data.
func Int64() (int64, error) {
	i, err := Uint64()

	if err != nil {
		return int64(0), err
	}

	return int64(i), nil
}

// Uint64 returns a random 64-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func Uint64() (uint64, error) {
	var bytes [8]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint64(0), err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}

/* __________________________________________________ */
