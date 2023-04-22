package command

import "github.com/funtimecoding/go-library/object_notation"

func (c *Command) Size() int {
	return len(object_notation.Encode(c, false))
}
