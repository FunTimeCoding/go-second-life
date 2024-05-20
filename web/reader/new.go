package reader

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func New(
	h *sentry.Hub,
	r *http.Request,
) *Reader {
	return &Reader{hub: h, request: r}
}
