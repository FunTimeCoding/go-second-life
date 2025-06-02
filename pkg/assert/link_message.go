package assert

import (
	"github.com/funtimecoding/go-second-life/pkg/link_message"
	"testing"
)

func LinkMessage(
	t *testing.T,
	expected link_message.Message,
	actual link_message.Message,
) {
	t.Helper()

	if actual.N == expected.N {
		t.Logf("PASS: Link event: %d", actual.N)
	} else {
		t.Errorf(
			"FAIL: Number"+
				"\nExpected: %d"+
				"\nActual:   %d",
			expected.N,
			actual.N,
		)
	}

	if actual.I == expected.I {
		t.Logf("PASS: UUID: '%s'", actual.I)
	} else {
		t.Errorf(
			"FAIL: Link identifier"+
				"\nExpected: %s"+
				"\nActual:   %s",
			expected.I,
			actual.I,
		)
	}

	if actual.T == expected.T {
		t.Logf("PASS: Text: '%s'", actual.T)
	} else {
		t.Errorf(
			"FAIL: Link text"+
				"\nExpected: %s"+
				"\nActual:   %s",
			expected.T,
			actual.T,
		)
	}
}
