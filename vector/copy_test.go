package vector

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	a := New(1, 0, 0)
	b := Copy(a)
	assert.Any(t, &Vector{X: 1}, b)
}
