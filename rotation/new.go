package rotation

func New(
	x float64,
	y float64,
	z float64,
	s float64,
) *Rotation {
	return &Rotation{X: x, Y: y, Z: z, S: s}
}
