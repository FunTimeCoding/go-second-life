package command

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestEscape(t *testing.T) {
	assert.String(t, Escape("a,b,c"), "a|b|c")
}
