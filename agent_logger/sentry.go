package agent_logger

func (l *Logger) Sentry(e error) {
	if e == nil {
		return
	}

	l.Sentryf("%s", e.Error())
}
