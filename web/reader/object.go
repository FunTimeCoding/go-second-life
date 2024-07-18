package reader

import "github.com/funtimecoding/go-library/pkg/notation"

func (r *Reader) Object(structure any) {
	notation.DecodeStrict(r.Text(), structure, false)
}
