package writer

import (
	"github.com/funtimecoding/go-library/pkg/integers"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web"
	secondLifeWeb "github.com/funtimecoding/go-second-life/pkg/web"
	"github.com/getsentry/sentry-go"
	"log"
)

func (w *Writer) Object(structure any) {
	encoded := notation.Encode(structure, false)
	web.ObjectHeader(w.writer)
	l := len(encoded)
	w.hub.ConfigureScope(
		func(s *sentry.Scope) {
			s.SetTag("response.body_length", integers.ToString(l))
		},
	)

	if l > secondLifeWeb.MaximumResponseSize {
		log.Printf("body too large: %d", l)
	} else if l > secondLifeWeb.LargeResponseThreshold {
		log.Printf("large body: %d", l)
	} else if l > secondLifeWeb.MediumResponseThreshold {
		log.Printf("medium body: %d", l)
	}

	w.bytes([]byte(encoded))
}
