package dialog

import (
	"github.com/funtimecoding/go-second-life/pkg/alias"
	"github.com/funtimecoding/go-second-life/pkg/alias/button"
	"golang.org/x/exp/slices"
	"sort"
)

func SortAndPadButtons(b []alias.Button) []alias.Button {
	sort.Slice(
		b,
		func(
			i int,
			j int,
		) bool {
			return b[i] < b[j]
		},
	)
	length := len(b)
	var toPad int

	if length <= 12 {
		slices.Reverse(b)
		// Pad to 12
		toPad = 12 - length
	} else if length%10 == 0 {
		toPad = 0
	} else {
		// Pad to the next 10
		modulo := length % 10
		toPad = 10 - modulo
	}

	var pad []alias.Button

	for i := 0; i < toPad; i++ {
		pad = append(pad, padButton)
	}

	if length <= 12 {
		// Index order without previous- and next-buttons:
		// 9 10 11
		// 6  7  8
		// 3  4  5
		// 0  1  2
		// Strategy:
		// 1. order list alphabetically in reverse
		// 2. pad if necessary
		// 3. flip left column with right column, row by row
		b = append(pad, b...)

		for i := 0; i < 12; i += 3 {
			button.Swap(b, i, i+2)
		}
	} else {
		// Index order with previous- and next-buttons:
		// 7 8 9
		// 4 5 6
		// 1 2 3
		// P 0 N
		// Strategy:
		// 1. order list alphabetically
		// 2. pad if necessary
		// 3.1 alphabetically reverse blocks of 10
		// 3.2 flip left column with right column, row by row
		b = append(b, pad...)
		newLength := len(b)

		for i := 0; i < newLength; i += 10 {
			slices.Reverse(b[i : i+10])
			button.Swap(b, i+7, i+9)
			button.Swap(b, i+4, i+6)
			button.Swap(b, i+1, i+3)
		}
	}

	return b
}
