package command

import "github.com/funtimecoding/go-second-life/link_message"

func (c *Command) AddLink(
	m link_message.Message,
	queued bool,
) {
	if queued {
		c.Q = append(c.Q, m)
	} else {
		c.L = append(c.L, m)
	}
}
