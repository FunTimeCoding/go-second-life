package command

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func Escape(s string) string {
	return strings.ReplaceAll(s, separator.Comma, separator.Pipe)
}
