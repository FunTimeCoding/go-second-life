package command

import "github.com/funtimecoding/go-second-life/link_message"

type Command struct {
	// L Link messages
	L []link_message.Message
	// O Owner say
	O string
}
