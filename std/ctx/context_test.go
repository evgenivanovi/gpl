package ctx

import (
	"context"
	"testing"
)

func TestDetachedContext_WithValue(t *testing.T) {
	t.Run("WithValueDetached", func(t *testing.T) {
		// given
		parent := context.TODO()

		// when
		detached := WithValueDetached(parent, "TEST_KEY", "TEST_VALUE")

		// then
		if value := parent.Value("TEST_KEY"); value != nil {
			t.Errorf(
				"Parent context contains key: '%s'",
				"TEST_KEY",
			)
		}

		if value := detached.Value("TEST_KEY").(string); "TEST_VALUE" != value {
			t.Errorf(
				"Actual value from context '%s' != expected: '%s'",
				value,
				"TEST_VALUE",
			)
		}
	})
}
