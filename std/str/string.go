package str

import (
	"strings"

	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

func Join(first string, second ...string) string {
	result := first
	for _, val := range second {
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
