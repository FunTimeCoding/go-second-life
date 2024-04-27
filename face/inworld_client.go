package face

import "github.com/funtimecoding/go-second-life/command"

type InworldClient interface {
	IsInworldClient()

	Post(
		locator string,
		o *command.Command,
	)
}
