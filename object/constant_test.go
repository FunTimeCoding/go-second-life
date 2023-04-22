package object

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, -1, AllSides)
}
