package agent_logger

import (
	"fmt"
	"github.com/funtimecoding/go-second-life/pkg/key"
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
		e = fmt.Errorf("%s", t)
	} else {
		e = fmt.Errorf("%s %s", l.agent, t)
	}

	l.hub.CaptureException(e)
}
