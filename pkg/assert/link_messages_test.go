package assert

import (
	"github.com/funtimecoding/go-second-life/pkg/link_message"
	"testing"
)

func TestLinkMessages(t *testing.T) {
	LinkMessages(
		t,
		nil,
		[]link_message.Message{
			link_message.New(1, "a"),
			link_message.New(2, "b"),
		},
	)
}
