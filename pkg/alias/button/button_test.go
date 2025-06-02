package button

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-second-life/pkg/alias"
	"testing"
)

func TestButton(t *testing.T) {
	assert.Any(t, []alias.Button{"Alfa"}, Slice([]string{strings.Alfa}))
}
