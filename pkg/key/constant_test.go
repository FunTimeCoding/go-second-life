package key

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(
		t,
		"00000000-0000-0000-0000-000000000000",
		Null,
	)
}
