package assert

import (
	"github.com/funtimecoding/go-second-life/pkg/command"
	"github.com/funtimecoding/go-second-life/pkg/link_message"
	"testing"
)

func TestCommand(t *testing.T) {
	c := command.New()
	c.AddLink(link_message.New(1, "a"))
	Command(
		t,
		&command.Command{L: []link_message.Message{{N: 1, I: "a"}}},
		c,
	)
}
