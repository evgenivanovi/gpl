package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinWithSep(t *testing.T) {
	tests := []struct {
		name     string
		first    string
		sep      string
		others   []string
		expected string
	}{
		{
			name:     "No first, no others, no separator",
			first:    "",
			sep:      "",
			others:   []string{},
			expected: "",
		},
		{
			name:     "With first, no others, no separator",
			first:    "first",
			sep:      "",
			others:   []string{},
			expected: "first",
		},
		{
			name:     "With first, no others, with separator",
			first:    "first",
			sep:      ";",
			others:   []string{},
			expected: "first",
		},
		{
			name:     "With first, with other, no separator",
			first:    "first",
			sep:      "",
			others:   []string{"second"},
			expected: "firstsecond",
		},
		{
			name:     "With first, with other, with separator",
			first:    "first",
			sep:      ",",
			others:   []string{"second"},
			expected: "first,second",
		},
		{
			name:     "With first, with others, no separator",
			first:    "first",
			sep:      "",
			others:   []string{"second", "third"},
			expected: "firstsecondthird",
		},
		{
			name:     "With first, with others, with separator",
			first:    "first",
			sep:      ",",
			others:   []string{"second", "third"},
			expected: "first,second,third",
		},
		{
			name:     "No first, with others, with separator",
			first:    "",
			sep:      ",",
			others:   []string{"second", "third"},
			expected: "second,third",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			actual := JoinWithSep(tt.first, tt.sep, tt.others...)

			// then
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sub      string
		expected string
	}{
		{
			name:     "Empty input, empty sub",
			input:    "",
			sub:      "",
			expected: "",
		},
		{
			name:     "Empty input",
			input:    "",
			sub:      "foo",
			expected: "",
		},
		{
			name:     "Empty sub",
			input:    "foo",
			sub:      "",
			expected: "foo",
		},
		{
			name:     "No occurrence",
			input:    "foo",
			sub:      "bar",
			expected: "foo",
		},
		{
			name:     "Single occurrence",
			input:    "foobar",
			sub:      "bar",
			expected: "foo",
		},
		{
			name:     "Multiple occurrences",
			input:    "foobarfoobar",
			sub:      "bar",
			expected: "foofoobar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			actual := RemoveFirst(tt.input, tt.sub)

			// then
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sub      string
		expected string
	}{
		{
			name:     "Empty input, empty sub",
			input:    "",
			sub:      "",
			expected: "",
		},
		{
			name:     "Empty input",
			input:    "",
			sub:      "foo",
			expected: "",
		},
		{
			name:     "Empty sub",
			input:    "foo",
			sub:      "",
			expected: "foo",
		},
		{
			name:     "No occurrence",
			input:    "foo",
			sub:      "bar",
			expected: "foo",
		},
		{
			name:     "Single occurrence",
			input:    "foobar",
			sub:      "bar",
			expected: "foo",
		},
		{
			name:     "Multiple occurrences",
			input:    "foobarfoobar",
			sub:      "bar",
			expected: "foofoo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			actual := RemoveAll(tt.input, tt.sub)

			// then
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTruncateAfter(t *testing.T) {

	// given
	tests := []struct {
		input     string
		delimiter string
		expected  string
	}{
		{
			"",
			"",
			"",
		},
		{
			"",
			"foo",
			"",
		},
		{
			"foo",
			"",
			"foo",
		},
		{
			"foo",
			"foo",
			"",
		},
		{
			"foofoo",
			"foo",
			"",
		},
	}

	for _, test := range tests {

		t.Run("Truncate after", func(t *testing.T) {

			// when
			actual := TruncateAfter(test.input, test.delimiter)

			// then
			assert.Equal(t, test.expected, actual)

		})

	}

}
