package command

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	c := New()
	assert.Any(t, &Command{}, c)
}
