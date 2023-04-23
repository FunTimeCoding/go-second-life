package command

func (c *Command) RemoveLastMessage() bool {
	position := len(c.Q) - 1

	if position == -1 || position == 0 {
		c.Q = nil

		return true
	}

	c.Q = c.Q[:position]

	return false
}
