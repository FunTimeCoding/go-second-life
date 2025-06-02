package agent_logger

import (
	"fmt"
	"github.com/funtimecoding/go-second-life/pkg/key"
)

func (l *Logger) Agent(
	s string,
	arguments ...any,
) {
	var t string

	if len(arguments) > 0 {
		t = fmt.Sprintf(s, arguments...)
	} else {
		t = s
	}

	if l.prefix != "" {
		t = fmt.Sprintf("%s: %s", l.prefix, t)
	}

	if l.agent != key.Null {
		t = fmt.Sprintf("%s %s", l.agent, t)
	}

	l.generic.Print(t)
}
