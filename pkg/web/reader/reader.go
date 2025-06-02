package reader

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

type Reader struct {
	hub     *sentry.Hub
	request *http.Request
}
