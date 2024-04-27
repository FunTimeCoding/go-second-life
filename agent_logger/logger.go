package agent_logger

import (
	"github.com/funtimecoding/go-second-life/key"
	"log"
)

type Logger struct {
	generic *log.Logger
	agent   key.Key
	prefix  string
}
