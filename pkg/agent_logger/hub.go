package agent_logger

import "github.com/getsentry/sentry-go"

func (l *Logger) Hub() *sentry.Hub {
	return l.hub
}
