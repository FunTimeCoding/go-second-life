package particle

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-second-life/color"
	"github.com/funtimecoding/go-second-life/vector"
	"math"
	"testing"
)

func TestList(t *testing.T) {
	assert.Any(
		t,
		[]string{
			"0",
			"1",
			"1",
			"<1,0,0>",
			"3",
			"<1,1,0>",
			"2",
			"0.2",
			"4",
			"0.7",
			"5",
			"<0.35,0.45,0.55>",
			"6",
			"<0.65,0.75,0.85>",
			"7",
			"2",
			"8",
			"<0.01,0.02,0.03>",
			"9",
			"3",
			"12",
			"texture",
			"13",
			"4",
			"15",
			"5",
			"16",
			"6",
			"17",
			"7",
			"18",
			"8",
			"19",
			"9",
			"20",
			"target",
			"21",
			"<0,0,0>",
			"22",
			"1.5707964",
			"23",
			"3.1415927",
		},
		List(
			1,
			color.New(1, 0, 0),
			color.New(1, 1, 0),
			0.2,
			0.7,
			vector.New(0.35, 0.45, 0.55),
			vector.New(0.65, 0.75, 0.85),
			2,
			vector.New(0.01, 0.02, 0.03),
			3,
			"texture",
			4,
			5,
			6,
			7,
			8,
			9,
			"target",
			vector.Zero,
			math.Pi/2,
			math.Pi,
		),
	)
}
