package vector

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestRoundVector(t *testing.T) {
	allSet := New(1.234, 2.345, 3.456)
	assertRound(t, allSet, "<1.234,2.345,3.456>")
	assertRound(t, allSet.Round(1), "<1.2,2.3,3.5>")
}

func assertRound(
	t *testing.T,
	vector *Vector,
	expected string,
) {
	t.Helper()
	assert.String(t, vector.String(), expected)
}
