package button

import "github.com/funtimecoding/go-second-life/alias"

func Slice(v []string) []alias.Button {
	var result []alias.Button

	for _, element := range v {
		result = append(result, alias.Button(element))
	}

	return result
}
