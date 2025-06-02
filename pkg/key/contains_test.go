package key

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestContains(t *testing.T) {
	assert.True(
		t,
		Contains(
			[]Key{New(strings.Alfa), New(strings.Bravo)},
			New(strings.Alfa),
		),
	)
	assert.False(
		t,
		Contains(
			[]Key{New(strings.Alfa), New(strings.Bravo)},
			New(strings.Charlie),
		),
	)
}
