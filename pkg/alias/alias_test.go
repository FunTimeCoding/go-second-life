package alias

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAlias(t *testing.T) {
	assert.String(t, "button", Button("button").String())
	assert.String(t, "menu", Menu("menu").String())
}
