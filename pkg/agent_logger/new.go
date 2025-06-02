package agent_logger

import (
	"github.com/funtimecoding/go-second-life/pkg/key"
	"github.com/getsentry/sentry-go"
	"log"
)

func New(
	l *log.Logger,
	h *sentry.Hub,
	agent key.Key,
) *Logger {
	return &Logger{generic: l, hub: h, agent: agent}
}
