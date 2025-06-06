package sentry

import (
	"github.com/funtimecoding/go-second-life/pkg/vector"
	"github.com/getsentry/sentry-go"
)

func SetRegionDetails(
	h *sentry.Hub,
	regionName string,
	regionVersion string,
	position *vector.Vector,
) {
	h.ConfigureScope(
		func(s *sentry.Scope) {
			s.SetTag("region", regionName)
			s.SetTag("region.version", regionVersion)
			s.SetExtra("position", position.String())
		},
	)
}
