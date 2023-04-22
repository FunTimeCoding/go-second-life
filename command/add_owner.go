package command

import (
	"fmt"
	"strings"
)

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

	text = strings.ReplaceAll(text, ",", Pipe)

	if c.O == "" {
		c.O = text
	} else {
		c.O += "\n" + text
	}
}
