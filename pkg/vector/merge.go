package vector

import "github.com/funtimecoding/go-library/pkg/math"

func (v *Vector) Merge(
	other *Vector,
	ownWeight float64,
	otherWeight float64,
) *Vector {
	var newVector Vector
	newVector.X = math.Weight(v.X, other.X, ownWeight, otherWeight)
	newVector.Y = math.Weight(v.Y, other.Y, ownWeight, otherWeight)
	newVector.Z = math.Weight(v.Z, other.Z, ownWeight, otherWeight)

	return &newVector
}
