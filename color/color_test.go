package color

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestNew(t *testing.T) {
	// minimum
	assert.Any(t, &Color{R: 0, G: 0, B: 0}, New(0, 0, 0))

	grey := New(0.5, 0.5, 0.5)
	assert.Round(t, 0.5, grey.R, 6)
	assert.Round(t, 0.5, grey.G, 6)
	assert.Round(t, 0.5, grey.B, 6)

	// limit
	assert.Any(t, &Color{R: 1, G: 1, B: 1}, New(1, 1, 1))

	// exceed limit
	assert.Any(t, &Color{R: 1, G: 1, B: 1}, New(2, 2, 2))
}

func TestNewRGB(t *testing.T) {
	// minimum
	assert.Any(t, &Color{R: 0, G: 0, B: 0}, NewRGB(0, 0, 0))

	grey := NewRGB(128, 128, 128)
	assert.Round(t, 0.501961, grey.R, 6)
	assert.Round(t, 0.501961, grey.G, 6)
	assert.Round(t, 0.501961, grey.B, 6)

	// limit
	assert.Any(
		t, &Color{R: 1, G: 1, B: 1},
		NewRGB(255, 255, 255),
	)

	// exceed limit
	assert.Any(
		t, &Color{R: 1, G: 1, B: 1},
		NewRGB(256, 256, 256),
	)
}

func TestToString(t *testing.T) {
	assert.String(t, "<0,0,0>", New(0, 0, 0).String())
	assert.String(
		t, "<128,128,128>",
		FromStringRGB("<128,128,128>").StringRGB(),
	)
}

func TestToStringRGB(t *testing.T) {
	assert.String(
		t, "<0,0,0>",
		New(0, 0, 0).StringRGB(),
	)
}

func TestFromString(t *testing.T) {
	// minimum
	assert.Any(t, &Color{R: 0, G: 0, B: 0}, FromString("<0,0,0>"))

	// gray
	grey := FromString("<0.5,0.5,0.5>")
	assert.Round(t, 0.5, grey.R, 6)
	assert.Round(t, 0.5, grey.G, 6)
	assert.Round(t, 0.5, grey.B, 6)

	// limit
	assert.Any(t, &Color{R: 1, G: 1, B: 1}, FromString("<1,1,1>"))

	// exceed limit
	assert.Any(t, &Color{R: 1, G: 1, B: 1}, FromString("<2,2,2>"))
}

func TestFromStringRGB(t *testing.T) {
	// minimum
	assert.Any(t, &Color{R: 0, G: 0, B: 0}, FromStringRGB("<0,0,0>"))

	// gray
	grey := FromStringRGB("<128,128,128>")
	assert.Round(t, 0.501961, grey.R, 6)
	assert.Round(t, 0.501961, grey.G, 6)
	assert.Round(t, 0.501961, grey.B, 6)

	// limit
	assert.Any(
		t, &Color{R: 1, G: 1, B: 1},
		FromStringRGB("<255,255,255>"),
	)

	// exceed limit
	assert.Any(
		t, &Color{R: 1, G: 1, B: 1},
		FromStringRGB("<256,256,256>"),
	)
}
