package sentry

import (
	"github.com/funtimecoding/go-second-life/pkg/key"
	"github.com/getsentry/sentry-go"
)

func SetObjectDetails(
	h *sentry.Hub,
	owner key.Key,
	ownerName string,
	object key.Key,
) {
	h.ConfigureScope(
		func(s *sentry.Scope) {
			s.SetUser(sentry.User{ID: owner.String(), Username: ownerName})
			s.SetTag("object.key", object.String())
		},
	)
}
