package agent_logger

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-second-life/key"
	"testing"
)

func TestLogger(t *testing.T) {
	assert.True(t, New(nil, nil, key.New(strings.Alfa)) != nil)
}
