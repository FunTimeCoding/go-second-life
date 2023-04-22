package color

import (
	"fmt"
	"github.com/funtimecoding/go-library/integers"
)

func (c *Color) StringRGB() string {
	return fmt.Sprintf(
		"<%s,%s,%s>",
		integers.ToString(int(c.R*float64(maximumInteger))),
		integers.ToString(int(c.G*float64(maximumInteger))),
		integers.ToString(int(c.B*float64(maximumInteger))),
	)
}
