package command

import "github.com/funtimecoding/go-second-life/link_message"

type Command struct {
	// Q Queued link messages
	Q []link_message.Message
	// O Owner say
	O string
}
