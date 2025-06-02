package command

import "fmt"

func (c *Command) AddOwner(
	format string,
	arguments ...any,
) {
	var text string

	if len(arguments) > 0 {
		text = fmt.Sprintf(format, arguments...)
	} else {
		text = format
	}

	text = Escape(text)

	if c.O == "" {
		c.O = text
	} else {
		c.O += "\n" + text
	}
}
