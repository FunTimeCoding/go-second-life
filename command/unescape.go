package command

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"strings"
)

func Unescape(s string) string {
	return strings.ReplaceAll(s, separator.Pipe, separator.Comma)
}
