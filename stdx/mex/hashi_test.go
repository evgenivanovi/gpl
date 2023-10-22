package mex

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAppendFormat(t *testing.T) {

	t.Parallel()

	t.Run("append format", func(t *testing.T) {

		// given
		input := []error{
			errors.New("FIRST"),
			errors.New("SECOND"),
			errors.New("THIRD"),
		}

		testee := AppendFormat("; ")

		// when
		actual := testee(input)

		// then
		expected := "FIRST; SECOND; THIRD"
		assert.Equal(t, expected, actual)

	})

}

func TestWithoutAfterSuffixFormat(t *testing.T) {

	t.Parallel()

	t.Run("append format", func(t *testing.T) {

		// given
		input := []error{
			errors.New("FIRST: err"),
			errors.New("SECOND: err"),
			errors.New("THIRD: err"),
		}

		testee := AppendWithoutAfterSuffixFormat("; ", ":")

		// when
		actual := testee(input)

		// then
		expected := "FIRST; SECOND; THIRD"
		assert.Equal(t, expected, actual)

	})

}
