package link_message

import "github.com/funtimecoding/go-library/pkg/separator"

func (m *Message) Arguments(a ...string) {
	m.ArgumentsSeparator(separator.Comma, a...)
}
