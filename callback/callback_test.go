package callback

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/notation"
	"testing"
)

func TestCallback(t *testing.T) {
	var c Callback
	notation.DecodeStrict(
		"{"+
			"\"Name\": \"Test\","+
			"\"Arguments\": [\"a\", \"text: 0.1\"]"+
			"}",
		&c,
	)
	assert.String(t, "Test", c.Name)
	assert.Any(t, []string{"a", "text: 0.1"}, c.Arguments)
}
