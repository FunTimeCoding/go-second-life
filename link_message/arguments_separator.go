package link_message

import "strings"

func (m *Message) ArgumentsSeparator(
	separator string,
	a ...string,
) {
	m.T = strings.Join(a, separator)
}
