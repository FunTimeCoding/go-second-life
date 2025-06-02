package vector

func New(
	x float64,
	y float64,
	z float64,
) *Vector {
	return &Vector{X: x, Y: y, Z: z}
}
