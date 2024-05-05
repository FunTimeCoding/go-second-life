package dialog

import "github.com/funtimecoding/go-library/pkg/strings"

const padButton = " "

func SortAndPadButtons(buttons []string) []string {
	length := len(buttons)
	var toPad int

	if length <= 12 {
		strings.Sort(buttons, false)
		// Pad to 12
		toPad = 12 - length
	} else if length%10 == 0 {
		strings.Sort(buttons, true)
		toPad = 0
	} else {
		strings.Sort(buttons, true)
		// Pad to the next 10
		modulo := length % 10
		toPad = 10 - modulo
	}

	var padding []string

	for i := 0; i < toPad; i++ {
		padding = append(padding, padButton)
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
		buttons = append(padding, buttons...)

		for i := 0; i < 12; i += 3 {
			strings.Swap(buttons, i, i+2)
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
		buttons = append(buttons, padding...)
		newLength := len(buttons)

		for i := 0; i < newLength; i += 10 {
			strings.Reverse(buttons[i : i+10])
			strings.Swap(buttons, i+7, i+9)
			strings.Swap(buttons, i+4, i+6)
			strings.Swap(buttons, i+1, i+3)
		}
	}

	return buttons
}
