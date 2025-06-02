package vector

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRegionPositionFromHeader(t *testing.T) {
	assert.Any(
		t,
		&Vector{X: 173, Y: 75.6, Z: 61},
		RegionPositionFromHeader(
			"(173.009827, 75.551231, 60.950001)",
		),
	)
}
