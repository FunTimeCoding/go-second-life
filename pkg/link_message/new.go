package link_message

import "github.com/funtimecoding/go-second-life/pkg/key"

func New(
	number int,
	k key.Key,
) Message {
	if k == key.Null {
		return Message{N: number}
	}

	return Message{N: number, I: k}
}
