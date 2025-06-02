package writer

import "github.com/funtimecoding/go-library/pkg/sentry"

func (w *Writer) bytes(body []byte) {
	_, e := w.writer.Write(body)
	sentry.CaptureOnError(w.hub, e)
}
