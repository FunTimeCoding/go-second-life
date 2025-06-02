package color

func NewRGB(
	red int,
	green int,
	blue int,
) *Color {
	if red < 0 {
		red = 0
	} else if red > maximumInteger {
		red = maximumInteger
	}

	if green < 0 {
		green = 0
	} else if green > maximumInteger {
		green = maximumInteger
	}

	if blue < 0 {
		blue = 0
	} else if blue > maximumInteger {
		blue = maximumInteger
	}

	return &Color{
		R: float64(red) / float64(maximumInteger),
		G: float64(green) / float64(maximumInteger),
		B: float64(blue) / float64(maximumInteger),
	}
}
