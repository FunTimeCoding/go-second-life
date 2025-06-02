package button

import "github.com/funtimecoding/go-second-life/pkg/alias"

func Swap(
	b []alias.Button,
	i int,
	j int,
) {
	b[i], b[j] = b[j], b[i]
}
