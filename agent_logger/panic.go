package agent_logger

func (l *Logger) Panic(e error) {
	if e != nil {
		l.Panicf(e.Error())
	}
}
