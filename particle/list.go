package particle

import (
	"github.com/funtimecoding/go-library/pkg/floats"
	"github.com/funtimecoding/go-library/pkg/integers"
	"github.com/funtimecoding/go-second-life/color"
	"github.com/funtimecoding/go-second-life/vector"
)

func List(
	flags int,
	startColor *color.Color,
	endColor *color.Color,
	startAlpha float64,
	endAlpha float64,
	startScale *vector.Vector,
	endScale *vector.Vector,
	age float64,
	acceleration *vector.Vector,
	pattern int,
	texture string,
	burstRate float64,
	count int,
	radius float64,
	minimumSpeed float64,
	maximumSpeed float64,
	maximumAge float64,
	target string,
	omega *vector.Vector,
	startAngle float64,
	endAngle float64,
) []string {
	return []string{
		integers.ToString(Flags),
		integers.ToString(flags),

		integers.ToString(StartColor),
		startColor.String(),

		integers.ToString(EndColor),
		endColor.String(),

		integers.ToString(StartAlpha),
		floats.ToString(startAlpha),

		integers.ToString(EndAlpha),
		floats.ToString(endAlpha),

		integers.ToString(StartScale),
		startScale.String(),

		integers.ToString(EndScale),
		endScale.String(),

		integers.ToString(MaximumAge),
		floats.ToString(age),

		integers.ToString(SourceAcceleration),
		acceleration.String(),

		integers.ToString(SourcePattern),
		integers.ToString(pattern),

		integers.ToString(SourceTexture),
		texture,

		integers.ToString(SourceBurstRate),
		floats.ToString(burstRate),

		integers.ToString(SourceBurstParticleCount),
		integers.ToString(count),

		integers.ToString(SourceBurstRadius),
		floats.ToString(radius),

		integers.ToString(SourceBurstSpeedMinimum),
		floats.ToString(minimumSpeed),

		integers.ToString(SourceBurstSpeedMaximum),
		floats.ToString(maximumSpeed),

		integers.ToString(SourceMaximumAge),
		floats.ToString(maximumAge),

		integers.ToString(SourceTargetKey),
		target,

		integers.ToString(SourceOmega),
		omega.String(),

		integers.ToString(SourceAngleBegin),
		floats.ToString(startAngle),

		integers.ToString(SourceAngleEnd),
		floats.ToString(endAngle),
	}
}
