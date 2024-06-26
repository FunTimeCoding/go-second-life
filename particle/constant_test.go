package particle

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Any(
		t,
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
		[]int{
			Flags,
			StartColor,
			StartAlpha,
			EndColor,
			EndAlpha,
			StartScale,
			EndScale,
			MaximumAge,
		},
	)
	assert.Any(
		t,
		[]int{8, 9, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23},
		[]int{
			SourceAcceleration,
			SourcePattern,
			SourceTexture,
			SourceBurstRate,
			SourceBurstParticleCount,
			SourceBurstRadius,
			SourceBurstSpeedMinimum,
			SourceBurstSpeedMaximum,
			SourceMaximumAge,
			SourceTargetKey,
			SourceOmega,
			SourceAngleBegin,
			SourceAngleEnd,
		},
	)
	assert.Any(
		t,
		[]int{1, 2, 4, 8, 16},
		[]int{
			PatternDrop,
			PatternExplode,
			PatternAngle,
			PatternAngleCone,
			PatternAngleConeEmpty,
		},
	)
	assert.Any(
		t,
		[]int{0, 1, 2, 4, 16, 32, 64, 128, 256, 1024},
		[]int{
			FlagOff,
			FlagInterpolateColorMask,
			FlagInterpolateScaleMask,
			FlagBounceMask,
			FlagFollowSourceMask,
			FlagFollowVelocityMask,
			FlagTargetPositionMask,
			FlagTargetLinearMask,
			FlagEmissiveMask,
			FlagRibbonMask,
		},
	)
	assert.Integer(
		t,
		259,
		FlagInterpolateColorMask|FlagInterpolateScaleMask|FlagEmissiveMask,
	)
}
