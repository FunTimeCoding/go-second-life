package color

import (
	stringsHelper "github.com/funtimecoding/go-library/strings"
	"strings"
)

func FromStringRGB(value string) *Color {
	elements := strings.Split(strings.Trim(value, "<>"), ",")

	return NewRGB(
		stringsHelper.ToInteger(elements[0], 0),
		stringsHelper.ToInteger(elements[1], 0),
		stringsHelper.ToInteger(elements[2], 0),
	)
}
