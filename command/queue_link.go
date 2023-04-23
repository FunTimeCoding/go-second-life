package command

import "github.com/funtimecoding/go-second-life/link_message"

func (c *Command) AddLink(m link_message.Message) {
	c.L = append(c.L, m)
}
