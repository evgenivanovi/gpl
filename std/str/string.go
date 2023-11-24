package str

import (
	"strings"

	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

func Join(first string, others ...string) string {
	result := first
	for _, val := range others {
		result += val
	}
	return result
}

func JoinWithSep(first string, sep string, others ...string) string {
	result := first
	for _, val := range others {
		result += sep
		result += val
	}
	return result
}

func RemoveFirst(value string, sub string) string {
	return strings.Replace(value, sub, std.Empty, 1)
}

func RemoveAll(value string, sub string) string {
	return strings.Replace(value, sub, std.Empty, -1)
}

/* __________________________________________________ */
