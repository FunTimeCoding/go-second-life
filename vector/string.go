package vector

import (
	"fmt"
	"github.com/funtimecoding/go-library/floats"
)

func (v *Vector) String() string {
	return fmt.Sprintf(
		"<%s,%s,%s>",
		floats.ToString(v.X),
		floats.ToString(v.Y),
		floats.ToString(v.Z),
	)
}
