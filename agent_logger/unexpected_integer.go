package agent_logger

func (l *Logger) UnexpectedInteger(v int) {
	l.Panicf("unexpected: %d", v)
}
