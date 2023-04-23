package command

import "strings"

func (c *Command) ReplaceLinkMessage(
	prefix string,
	replace string,
) {
	for i, element := range c.Q {
		if strings.HasPrefix(element.T, prefix) {
			c.Q[i].T = replace
		}
	}
}
