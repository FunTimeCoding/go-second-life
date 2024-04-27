package agent_logger

func (l *Logger) Sentry(e error) {
	if e != nil {
		l.Sentryf(e.Error())
	}
}
