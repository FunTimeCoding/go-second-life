package input

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-second-life/alias"
	"github.com/funtimecoding/go-second-life/constant"
	"testing"
)

func TestInput(t *testing.T) {
	i := &Input{
		ButtonClicked: alias.Button(
			fmt.Sprintf(
				"%s%s",
				constant.TextPrefix,
				strings.Alfa,
			),
		),
		TextEntered: fmt.Sprintf(
			"%s%s",
			constant.TextPrefix,
			strings.Bravo,
		),
		ObjectTouched: fmt.Sprintf(
			"%s%s",
			constant.TextPrefix,
			strings.Charlie,
		),
	}
	i.TrimPrefixes()
	assert.Any(
		t,
		&Input{
			ButtonClicked: "Alfa",
			TextEntered:   "Bravo",
			ObjectTouched: "Charlie",
		},
		i,
	)
}
