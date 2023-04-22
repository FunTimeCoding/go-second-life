package rotation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestFromString(t *testing.T) {
	assert.Any(
		t,
		&Rotation{X: 1, Y: 2, Z: 3, S: 4},
		FromString("<1,2,3,4>"),
	)
}
