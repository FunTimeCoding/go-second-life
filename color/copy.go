package color

func Copy(v *Color) *Color {
	return New(v.R, v.G, v.B)
}
