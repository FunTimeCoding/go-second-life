package vector

import "math"

func (v *Vector) Distance(other *Vector) float64 {
	return math.Pow(
		math.Pow(other.X-v.X, 2)+
			math.Pow(other.Y-v.Y, 2)+
			math.Pow(other.Z-v.Z, 2),
		0.5,
	)
}
