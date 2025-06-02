package particle

// Flag
const (
	FlagOff                  = 0x0
	FlagInterpolateColorMask = 0x001
	FlagInterpolateScaleMask = 0x002
	FlagBounceMask           = 0x004
	FlagFollowSourceMask     = 0x010
	FlagFollowVelocityMask   = 0x020
	FlagTargetPositionMask   = 0x040
	FlagTargetLinearMask     = 0x080
	FlagEmissiveMask         = 0x100
	FlagRibbonMask           = 0x400
)

// Pattern
const (
	PatternDrop           = 0x01
	PatternExplode        = 0x02
	PatternAngle          = 0x04
	PatternAngleCone      = 0x08
	PatternAngleConeEmpty = 0x10
)

const (
	Flags      = int(0)
	StartColor = int(1)
	StartAlpha = int(2)
	EndColor   = int(3)
	EndAlpha   = int(4)
	StartScale = int(5)
	EndScale   = int(6)
	MaximumAge = int(7)
)

const (
	SourceAcceleration       = int(8)
	SourcePattern            = int(9)
	SourceTexture            = int(12)
	SourceBurstRate          = int(13)
	SourceBurstParticleCount = int(15)
	SourceBurstRadius        = int(16)
	SourceBurstSpeedMinimum  = int(17)
	SourceBurstSpeedMaximum  = int(18)
	SourceMaximumAge         = int(19)
	SourceTargetKey          = int(20)
	SourceOmega              = int(21)
	SourceAngleBegin         = int(22)
	SourceAngleEnd           = int(23)
)
