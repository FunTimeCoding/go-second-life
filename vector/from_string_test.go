package vector

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFromString(t *testing.T) {
	assert.Any(
		t,
		&Vector{X: 1, Y: 2, Z: 3},
		FromString("<1,2,3>"),
	)
}
