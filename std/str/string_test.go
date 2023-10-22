package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAfter(t *testing.T) {

	// given
	datas := []struct {
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
			"foobar",
			"bar",
			"foo",
		},
	}

	for _, data := range datas {

		t.Run("remove after", func(t *testing.T) {

			// when
			actual := RemoveAfter(data.input, data.delimiter)

			// then
			assert.Equal(t, data.expected, actual)

		})

	}

}
