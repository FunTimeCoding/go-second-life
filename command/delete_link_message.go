package command

func (c *Command) DeleteLinkMessage(
	number int,
	search []string,
) {
	index := c.SearchLinkMessage(number, search)

	if index != -1 {
		// Shift left
		copy(c.L[index:], c.L[index+1:])
		// Remove last
		c.L = c.L[:len(c.L)-1]
	}
}
