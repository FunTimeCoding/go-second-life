package link_message

import "github.com/funtimecoding/go-second-life/pkg/key"

type Message struct {
	// N Number
	N int
	// T Text
	T string
	// I Identifier
	I key.Key
}
