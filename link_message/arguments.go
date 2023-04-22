package link_message

func (m *Message) Arguments(a ...string) {
	m.ArgumentsSeparator(",", a...)
}
