package key

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestContainsAny(t *testing.T) {
	assert.True(
		t,
		ContainsAny(
			[]Key{New(strings.Alfa), New(strings.Bravo)},
			[]Key{New(strings.Alfa)},
		),
	)
	assert.False(
		t,
		ContainsAny(
			[]Key{New(strings.Alfa), New(strings.Bravo)},
			[]Key{New(strings.Charlie)},
		),
	)
	assert.False(
		t,
		ContainsAny(
			[]Key{New(strings.Alfa), New(strings.Bravo)},
			[]Key{},
		),
	)
	assert.False(
		t,
		ContainsAny(
			[]Key{},
			[]Key{New(strings.Alfa)},
		),
	)
}
