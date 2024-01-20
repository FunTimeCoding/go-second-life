package assert

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-second-life/command"
	"testing"
)

func Command(
	t *testing.T,
	expected *command.Command,
	actual *command.Command,
) {
	t.Helper()
	t.Logf("%s", notation.Format(actual))
	expectedCount := len(expected.L)
	actualCount := len(actual.L)

	if actualCount != expectedCount {
		t.Errorf(
			"Message count"+
				"\nExpected: %d"+
				"\nActual:   %d",
			expectedCount,
			actualCount,
		)
	}

	LinkMessages(t, expected.L, actual.L)

	if actual.O != expected.O {
		t.Errorf(
			"Owner say"+
				"\nExpected: %s"+
				"\nActual:   %s",
			expected.O,
			actual.O,
		)
	}
}
