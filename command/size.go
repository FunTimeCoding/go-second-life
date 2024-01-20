package command

import "github.com/funtimecoding/go-library/pkg/notation"

func (c *Command) Size() int {
	return len(notation.Encode(c, false))
}
