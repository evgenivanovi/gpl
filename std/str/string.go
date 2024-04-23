package str

import (
	"strings"

	"github.com/evgenivanovi/gpl/std"
)

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

func RemoveAfter(
	input string,
	delimiter string,
) string {
	return RemoveAfterWithMissing(
		input, delimiter, input,
	)
}

func RemoveAfterWithMissing(
	input string,
	delimiter string,
	missingDelimiterValue string,
) string {

	if len(input) == 0 {
		return input
	}

	parts := strings.SplitN(input, delimiter, 2)

	if len(parts) == 1 {
		return missingDelimiterValue
	}

	res := strings.Builder{}
	res.WriteString(parts[0])
	res.WriteString(parts[1][len(parts[1]):])

	return res.String()

}
