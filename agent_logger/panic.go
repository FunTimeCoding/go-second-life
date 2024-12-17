package agent_logger

func (l *Logger) Panic(e error) {
	if e == nil {
		return
	}

	l.Panicf("%s", e.Error())
}
