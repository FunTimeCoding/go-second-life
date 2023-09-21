package key

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestKey(t *testing.T) {
	k := New("a")
	assert.String(t, "a", k.String())
}
