package str

import (
	"strings"

	"github.com/evgenivanovi/gpl/std"
)

// Join concatenates the first string with any additional strings provided.
func Join(first string, others ...string) string {
	return JoinWithSep(first, std.Empty, others...)
}

// JoinWithSep concatenates the first string with any additional strings provided,
// Uses specified separator between each string.
func JoinWithSep(first string, sep string, others ...string) string {
	result := strings.Builder{}
	result.WriteString(first)

	for ind := range others {
		if result.Len() > std.Zero {
			result.WriteString(sep)
		}
		result.WriteString(others[ind])
	}

	return result.String()
}

func RemoveFirst(value string, sub string) string {
	return strings.Replace(value, sub, std.Empty, std.One)
}

func RemoveAll(value string, sub string) string {
	return strings.Replace(value, sub, std.Empty, std.MinusOne)
}

// TruncateAfter returns substring of `value`
// that appears before the first occurrence of `delimiter`.
// If `delimiter` was not found or empty, `value` is returned unchanged.
func TruncateAfter(value string, delimiter string) string {
	if len(value) == 0 || len(delimiter) == 0 {
		return value
	}
	return strings.SplitN(value, delimiter, 2)[0]
}
