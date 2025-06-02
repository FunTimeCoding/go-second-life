package command

func (c *Command) RemoveLast() bool {
	position := len(c.L) - 1

	if position == -1 || position == 0 {
		c.L = nil

		return true
	}

	c.L = c.L[:position]

	return false
}
