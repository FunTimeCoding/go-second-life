package inventory

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPath(t *testing.T) {
	assert.String(t, "/test", Path("test"))
}
