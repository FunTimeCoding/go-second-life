package chat

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestAgentLocator(t *testing.T) {
	assert.String(
		t,
		"secondlife:///app/agent/a/about",
		AgentLocator("a"),
	)
}
