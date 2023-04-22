package rotation

import (
	stringsHelper "github.com/funtimecoding/go-library/strings"
	"strings"
)

func FromString(text string) *Rotation {
	parts := strings.Split(
		strings.TrimRight(strings.TrimLeft(text, "<"), ">"),
		",",
	)

	return New(
		stringsHelper.ToFloatStrict(parts[0]),
		stringsHelper.ToFloatStrict(parts[1]),
		stringsHelper.ToFloatStrict(parts[2]),
		stringsHelper.ToFloatStrict(parts[3]),
	)
}
