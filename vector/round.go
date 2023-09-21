package vector

import "github.com/funtimecoding/go-library/pkg/math/round"

func (v *Vector) Round(decimals int) *Vector {
	return New(
		round.Round(v.X, decimals),
		round.Round(v.Y, decimals),
		round.Round(v.Z, decimals),
	)
}
