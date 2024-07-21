package input

import (
	"github.com/funtimecoding/go-second-life/alias"
	"github.com/funtimecoding/go-second-life/constant"
	"strings"
)

// TrimPrefixes
// Remove text prefix used to prevent numbers from being cast which causes
// unmarshal to fail
func (i *Input) TrimPrefixes() {
	i.ButtonClicked = alias.Button(
		strings.TrimPrefix(i.ButtonClicked.String(), constant.TextPrefix),
	)
	i.TextEntered = strings.TrimPrefix(i.TextEntered, constant.TextPrefix)
	i.ObjectTouched = strings.TrimPrefix(i.ObjectTouched, constant.TextPrefix)
}
