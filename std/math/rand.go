package math

import (
	"crypto/rand"
	"encoding/binary"
)

// RandomInt8 returns a random 8-bit signed integer.
// Return 0 and an error if unable to get random data.
func RandomInt8() (int8, error) {
	i, err := RandomUint8()

	if err != nil {
		return int8(0), err
	}

	return int8(i), nil
}

// RandomUint8 returns a random 8-bit unsigned integer.
// Return 0 and an error if unable to get rand data.
func RandomUint8() (uint8, error) {
	var bytes [1]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint8(0), err
	}

	return bytes[0], nil
}

// RandomInt16 returns a random 16-bit signed integer.
// Return 0 and an error if unable to get rand data.
func RandomInt16() (int16, error) {
	i, err := RandomUint16()

	if err != nil {
		return int16(0), err
	}

	return int16(i), nil
}

// RandomUint16 returns a random 16-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func RandomUint16() (uint16, error) {
	var bytes [2]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint16(0), err
	}

	return binary.LittleEndian.Uint16(bytes[:]), nil
}

// RandomInt32 returns a random 32-bit signed integer.
// Return 0 and an error if unable to get random data.
func RandomInt32() (int32, error) {
	i, err := RandomUint32()

	if err != nil {
		return int32(0), err
	}

	return int32(i), nil
}

// RandomUint32 returns a rand 32-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func RandomUint32() (uint32, error) {
	var bytes [4]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint32(0), err
	}

	return binary.LittleEndian.Uint32(bytes[:]), nil
}

// RandomInt64 returns a random 64-bit signed integer.
// Return 0 and an error if unable to get random data.
func RandomInt64() (int64, error) {
	i, err := RandomUint64()

	if err != nil {
		return int64(0), err
	}

	return int64(i), nil
}

// RandomUint64 returns a random 64-bit unsigned integer.
// Return 0 and an error if unable to get random data.
func RandomUint64() (uint64, error) {
	var bytes [8]byte

	_, err := rand.Read(bytes[:])
	if err != nil {
		return uint64(0), err
	}

	return binary.LittleEndian.Uint64(bytes[:]), nil
}
