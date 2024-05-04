package agent_logger

import (
	"github.com/funtimecoding/go-second-life/key"
	"github.com/getsentry/sentry-go"
	"log"
)

type Logger struct {
	generic *log.Logger
	hub     *sentry.Hub
	agent   key.Key
	prefix  string
}
