package dtwire

import (
	"bytes"
	"fmt"
)

type TextPutter interface {
	PutText(w Writer) error
}

func PutText(w Writer, e Putter) error {
	if x, ok := e.(TextPutter); ok {
		return x.PutText(w)
	} else {
		return unknownTextPutter[Putter]{e}.PutText(w)
	}
}

type unknownTextPutter[T Putter] struct{ elem T }

func (p unknownTextPutter[T]) PutText(w Writer) error {
	size := p.elem.Size()
	var b bytes.Buffer
	b.Grow(int(size))
	p.elem.Put(&b)
	fmt.Fprintf(w, "<![CDATA[%x]]>\n", b.Bytes())
	return nil
}
