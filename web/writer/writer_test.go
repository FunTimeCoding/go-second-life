package writer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestWriter(t *testing.T) {
	assert.True(t, New(nil, nil) != nil)
}
