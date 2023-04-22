package rotation

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Any(t, &Rotation{}, Zero)
}
