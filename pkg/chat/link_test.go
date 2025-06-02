package chat

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"[secondlife:///app/chat/1/a b]",
		Link(1, "a", "b"),
	)
}
