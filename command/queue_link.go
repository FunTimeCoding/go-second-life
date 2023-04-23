package command

import "github.com/funtimecoding/go-second-life/link_message"

func (c *Command) QueueLink(m link_message.Message) {
	c.Q = append(c.Q, m)
}
