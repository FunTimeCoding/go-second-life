package particle

import (
	"github.com/funtimecoding/go-library/pkg/floats"
	"github.com/funtimecoding/go-library/pkg/integers"
	"github.com/funtimecoding/go-second-life/pkg/color"
	"github.com/funtimecoding/go-second-life/pkg/vector"
)

func List(
	flags int,
	startColor *color.Color,
	endColor *color.Color,
	startAlpha float64,
	endAlpha float64,
	startScale *vector.Vector,
	endScale *vector.Vector,
	maximumAge float64,
	sourceAcceleration *vector.Vector,
	sourcePattern int,
	sourceTexture string,
	sourceBurstRate float64,
	sourceCount int,
	sourceRadius float64,
	sourceMinimumSpeed float64,
	sourceMaximumSpeed float64,
	sourceMaximumAge float64,
	sourceTarget string,
	sourceOmega *vector.Vector,
	sourceStartAngle float64,
	sourceEndAngle float64,
) []string {
	return []string{
		integers.ToString(Flags), integers.ToString(flags),

		integers.ToString(StartColor), startColor.String(),

		integers.ToString(EndColor), endColor.String(),

		integers.ToString(StartAlpha), floats.ToString(startAlpha),

		integers.ToString(EndAlpha), floats.ToString(endAlpha),

		integers.ToString(StartScale), startScale.String(),

		integers.ToString(EndScale), endScale.String(),

		integers.ToString(MaximumAge), floats.ToString(maximumAge),

		integers.ToString(SourceAcceleration), sourceAcceleration.String(),

		integers.ToString(SourcePattern), integers.ToString(sourcePattern),

		integers.ToString(SourceTexture), sourceTexture,

		integers.ToString(SourceBurstRate), floats.ToString(sourceBurstRate),

		integers.ToString(SourceBurstParticleCount),
		integers.ToString(sourceCount),

		integers.ToString(SourceBurstRadius), floats.ToString(sourceRadius),

		integers.ToString(SourceBurstSpeedMinimum),
		floats.ToString(sourceMinimumSpeed),

		integers.ToString(SourceBurstSpeedMaximum),
		floats.ToString(sourceMaximumSpeed),

		integers.ToString(SourceMaximumAge),
		floats.ToString(sourceMaximumAge),

		integers.ToString(SourceTargetKey), sourceTarget,

		integers.ToString(SourceOmega), sourceOmega.String(),

		integers.ToString(SourceAngleBegin),
		floats.ToString(sourceStartAngle),

		integers.ToString(SourceAngleEnd),
		floats.ToString(sourceEndAngle),
	}
}
