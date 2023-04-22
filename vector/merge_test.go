package vector

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	assertMerge(
		t,
		&Vector{X: 1, Y: 1, Z: 1},
		&Vector{X: 1, Y: 1, Z: 1},
		1,
		1,
		&Vector{X: 1, Y: 1, Z: 1},
	)
	assertMerge(
		t,
		&Vector{X: 1, Y: 1, Z: 1},
		&Vector{X: 0.5, Y: 0.5, Z: 0.5},
		1,
		1,
		&Vector{X: 0.75, Y: 0.75, Z: 0.75},
	)
	assertMerge(
		t,
		&Vector{X: 256, Y: 128, Z: 64},
		&Vector{X: 128, Y: 64, Z: 32},
		2,
		1,
		&Vector{X: 213.33, Y: 106.67, Z: 53.33},
	)
}

func assertMerge(
	t *testing.T,
	a *Vector,
	b *Vector,
	aWeight float64,
	bWeight float64,
	expected *Vector,
) {
	t.Helper()
	assert.Any(t, expected, a.Merge(b, aWeight, bWeight).Round(2))
}
