package writer

func (w *Writer) Text(s string) {
	w.bytes([]byte(s))
}
