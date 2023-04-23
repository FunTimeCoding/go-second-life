package command

func (c *Command) DeleteLinkMessage(
	number int,
	search []string,
) {
	index := c.SearchLinkMessage(number, search)

	if index != -1 {
		// Shift left
		copy(c.Q[index:], c.Q[index+1:])
		// Remove last
		c.Q = c.Q[:len(c.Q)-1]
	}
}
