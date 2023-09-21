package rotation

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/floats"
)

func (r *Rotation) String() string {
	return fmt.Sprintf(
		"<%s,%s,%s,%s>",
		floats.ToString(r.X),
		floats.ToString(r.Y),
		floats.ToString(r.Z),
		floats.ToString(r.S),
	)
}
