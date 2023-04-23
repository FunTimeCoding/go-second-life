package command

func (c *Command) Empty() bool {
	if len(c.L) != 0 {
		return false
	}

	if c.O != "" {
		return false
	}

	return true
}
