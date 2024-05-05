package agent

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestFirstName(t *testing.T) {
	assertFirstName(
		t,
		"Test",
		"Test Something",
		"Test",
	)
	assertFirstName(
		t,
		"Test",
		"Test Resident",
		"Test",
	)
	assertFirstName(t, "", "Test", "Test")
	assertFirstName(t, "", "", "")
}

func assertFirstName(
	t *testing.T,
	display string,
	legacy string,
	expected string,
) {
	t.Helper()
	assert.String(t, expected, FirstName(display, legacy))
}
