package dialog

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestSortAndPadButtons(t *testing.T) {
	assert.Any(
		t,
		[]string{
			padButton, padButton, padButton,
			padButton, padButton, padButton,
			strings.Delta, padButton, padButton,
			strings.Alfa, strings.Bravo, strings.Charlie,
		},
		SortAndPadButtons(
			[]string{
				strings.Alfa,
				strings.Bravo,
				strings.Charlie,
				strings.Delta,
			},
		),
	)
	assert.Any(
		t,
		[]string{
			strings.Juliett, strings.Kilo, strings.Lima,
			strings.Golf, strings.Hotel, strings.India,
			strings.Delta, strings.Echo, strings.Foxtrot,
			strings.Alfa, strings.Bravo, strings.Charlie,
		},
		SortAndPadButtons(
			[]string{
				strings.Alfa, strings.Bravo, strings.Charlie,
				strings.Delta, strings.Echo, strings.Foxtrot,
				strings.Golf, strings.Hotel, strings.India,
				strings.Juliett, strings.Kilo, strings.Lima,
			},
		),
	)
	assert.Any(
		t,
		// First line is where < and > buttons are added by the Input module
		[]string{
			strings.Juliett,
			strings.Golf, strings.Hotel, strings.India,
			strings.Delta, strings.Echo, strings.Foxtrot,
			strings.Alfa, strings.Bravo, strings.Charlie,

			padButton,
			padButton, padButton, padButton,
			padButton, padButton, padButton,
			strings.Kilo, strings.Lima, strings.Mike,
		},
		SortAndPadButtons(
			[]string{
				strings.Alfa, strings.Bravo, strings.Charlie,
				strings.Delta, strings.Echo, strings.Foxtrot,
				strings.Golf, strings.Hotel, strings.India,
				strings.Juliett, strings.Kilo, strings.Lima,
				strings.Mike,
			},
		),
	)
}
