package agent_logger

import (
	"github.com/funtimecoding/go-second-life/key"
	"log"
)

func New(
	l *log.Logger,
	agent key.Key,
) *Logger {
	return &Logger{generic: l, agent: agent}
}
