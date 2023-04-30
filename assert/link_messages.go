package assert

import (
	"github.com/funtimecoding/go-second-life/link_message"
	"testing"
)

func LinkMessages(
	t *testing.T,
	expected []link_message.Message,
	actual []link_message.Message,
) {
	t.Helper()
	expectedLength := len(expected)
	actualLength := len(actual)

	for i, expectedElement := range expected {
		if i >= actualLength {
			t.Errorf(
				"Length mismatch"+
					"\nExpected: %d"+
					"\nActual: %d",
				expectedLength,
				actualLength,
			)

			break
		}

		LinkMessage(t, expectedElement, actual[i])
	}
}
