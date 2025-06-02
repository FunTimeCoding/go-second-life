package reader

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/getsentry/sentry-go"
	"io"
)

func (r *Reader) Text() string {
	body, e := io.ReadAll(r.request.Body)
	errors.PanicOnError(e)
	s := string(body)
	r.hub.ConfigureScope(
		func(s *sentry.Scope) {
			s.SetRequestBody(body)
			s.SetExtra("request_text", s)
		},
	)

	return s
}
