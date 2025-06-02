package color

import (
	stringsHelper "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"log"
	"strings"
)

func FromString(v string) *Color {
	s := strings.Split(strings.Trim(v, "<>"), separator.Comma)

	if len(s) != 3 {
		log.Panicf("Invalid color vector: %s %+v", v, s)
	}

	return New(
		stringsHelper.ToFloat(s[0], 0),
		stringsHelper.ToFloat(s[1], 0),
		stringsHelper.ToFloat(s[2], 0),
	)
}
