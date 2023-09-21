package vector

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestDistanceVector(t *testing.T) {
	assertDistance(t, New(0, 0, 0), New(1, 1, 1), 1.73)
}

func assertDistance(
	t *testing.T,
	vector *Vector,
	other *Vector,
	expected float64,
) {
	t.Helper()
	assert.Round(t, expected, vector.Distance(other), 2)
}
