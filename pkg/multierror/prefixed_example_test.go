package multierror

import (
	"bytes"
	"testing"
)

func TestExamplePrefixed(t *testing.T) {
	t.Run("testExample", func(t *testing.T) {
		var b = new(bytes.Buffer)
		output = b
		var want = new(bytes.Buffer)
		want.WriteString("config validation: 1 error occurred:\n\t* some validation error\n\n\n")

		ExamplePrefixed()

		if b.String() != want.String() {
			t.Errorf("ExamplePrefixed = %v, want = %v", b.String(), want.String())
		}
	})
}
