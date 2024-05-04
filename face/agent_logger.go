package face

import "github.com/getsentry/sentry-go"

type AgentLogger interface {
	IsAgentLogger()

	Prefix(s string)

	// Agent logs to standard output but does not panic
	Agent(
		s string,
		arguments ...any,
	)
	// Sentry logs to server but does not panic
	Sentry(e error)
	// Sentryf logs to server but does not panic
	Sentryf(
		s string,
		arguments ...any,
	)

	// Panic panics
	Panic(e error)
	// Panicf panics
	Panicf(
		s string,
		arguments ...any,
	)
	// Unexpected panics
	Unexpected(v string)
	// UnexpectedInteger panics
	UnexpectedInteger(v int)

	Hub() *sentry.Hub
}
