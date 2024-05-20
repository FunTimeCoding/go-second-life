package command

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestUnescape(t *testing.T) {
	assert.String(t, Unescape("a|b|c"), "a,b,c")
}
