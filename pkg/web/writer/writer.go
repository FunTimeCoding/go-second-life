package writer

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type Writer struct {
	hub    *sentry.Hub
	writer http.ResponseWriter
}
