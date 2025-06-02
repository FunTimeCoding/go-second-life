package rotation

import (
	stringsHelper "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func FromString(text string) *Rotation {
	parts := strings.Split(
		strings.TrimRight(strings.TrimLeft(text, "<"), ">"),
		separator.Comma,
	)

	return New(
		stringsHelper.ToFloatStrict(parts[0]),
		stringsHelper.ToFloatStrict(parts[1]),
		stringsHelper.ToFloatStrict(parts[2]),
		stringsHelper.ToFloatStrict(parts[3]),
	)
}
