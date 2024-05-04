package agent_logger

import (
	"fmt"
	"github.com/funtimecoding/go-second-life/key"
)

func (l *Logger) Sentryf(
	s string,
	arguments ...any,
) {
	var t string

	if len(arguments) > 0 {
		t = fmt.Sprintf(s, arguments...)
	} else {
		t = s
	}

	var e error

	if l.agent == key.Null {
		e = fmt.Errorf("%s\n", t)
	} else {
		e = fmt.Errorf("%s %s\n", l.agent, t)
	}

	l.hub.CaptureException(e)
}
