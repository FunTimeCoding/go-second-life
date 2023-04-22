package command

import "strings"

func (c *Command) ReplaceLinkMessage(
	prefix string,
	replace string,
) {
	for i, element := range c.L {
		if strings.HasPrefix(element.T, prefix) {
			c.L[i].T = replace
		}
	}
}
