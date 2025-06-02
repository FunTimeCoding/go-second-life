package writer

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func New(
	h *sentry.Hub,
	w http.ResponseWriter,
) *Writer {
	return &Writer{hub: h, writer: w}
}
