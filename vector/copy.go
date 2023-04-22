package vector

func Copy(
	v *Vector,
) *Vector {
	return New(v.X, v.Y, v.Z)
}
