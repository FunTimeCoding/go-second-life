package color

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	a := New(1, 0, 0)
	b := Copy(a)
	assert.Any(t, &Color{R: 1}, b)
}
