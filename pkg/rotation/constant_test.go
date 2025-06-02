package rotation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Any(t, &Rotation{}, Zero)
}
