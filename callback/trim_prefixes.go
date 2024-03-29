package callback

import (
	"github.com/funtimecoding/go-second-life/constant"
	"strings"
)

func (c *Callback) TrimPrefixes() {
	for i, element := range c.Arguments {
		c.Arguments[i] = strings.TrimPrefix(element, constant.TextPrefix)
	}
}
