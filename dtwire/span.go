package dtwire

import "fmt"

type Span[T Putter] struct {
	n    Uint64
	elem T
}

func SpanOf[T Putter](v T) Span[T] {
	return Span[T]{n: Uint64(v.Size()), elem: v}
}

func (x Span[T]) Elem() T { return x.elem }

func (x *Span[T]) Scan(r Reader) error {
	x.n.Scan(r)
	return any(&x.elem).(Scanner).Scan(r)
}

func (x Span[T]) Put(w Writer) error {
	x.n.Put(w)
	return x.elem.Put(w)
}

func (x Span[T]) Size() uint64 { return x.n.Size() + x.elem.Size() }

func (x Span[T]) PutText(w Writer) error {
	fmt.Fprintf(w, "<span size=%d>\n", x.n)
	PutText(w, x.elem)
	fmt.Fprintln(w, "</span>")
	return nil
}
