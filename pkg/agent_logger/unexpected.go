package agent_logger

func (l *Logger) Unexpected(v string) {
	l.Panicf("unexpected: %s", v)
}
