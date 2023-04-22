package color

func New(
	red float64,
	green float64,
	blue float64,
) *Color {
	if red < 0.0 {
		red = 0.0
	} else if red > 1.0 {
		red = 1.0
	}

	if green < 0.0 {
		green = 0.0
	} else if green > 1.0 {
		green = 1.0
	}

	if blue < 0.0 {
		blue = 0.0
	} else if blue > 1.0 {
		blue = 1.0
	}

	return &Color{
		R: red,
		G: green,
		B: blue,
	}
}
