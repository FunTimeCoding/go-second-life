package link_message

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestMessage(t *testing.T) {
	m := New(1, "a")
	assert.Integer(t, 1, m.N)
	assert.String(t, "a", m.I.String())
}
