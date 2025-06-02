package color

import (
	stringsHelper "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func FromStringRGB(value string) *Color {
	s := strings.Split(strings.Trim(value, "<>"), separator.Comma)

	return NewRGB(
		stringsHelper.ToInteger(s[0], 0),
		stringsHelper.ToInteger(s[1], 0),
		stringsHelper.ToInteger(s[2], 0),
	)
}
