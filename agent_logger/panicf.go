package agent_logger

import (
	"fmt"
	"github.com/funtimecoding/go-second-life/key"
)

func (l *Logger) Panicf(
	s string,
	arguments ...any,
) {
	var t string

	if len(arguments) > 0 {
		t = fmt.Sprintf(s, arguments...)
	} else {
		t = s
	}

	if l.agent == key.Null {
		l.generic.Panicf("%s\n", t)
	} else {
		l.generic.Panicf("%s %s\n", l.agent, t)
	}
}
