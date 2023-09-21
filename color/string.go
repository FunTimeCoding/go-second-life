package color

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/floats"
)

func (c *Color) String() string {
	return fmt.Sprintf(
		"<%s,%s,%s>",
		floats.ToString(c.R),
		floats.ToString(c.G),
		floats.ToString(c.B),
	)
}
