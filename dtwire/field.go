package dtwire

import "fmt"

type Field[T Putter] struct {
	num Uint64
	Span[T]
}

func FieldOf[T Putter](n uint64, v T) Field[T] {
	return Field[T]{num: Uint64(n), Span: SpanOf(v)}
}

func (f Field[T]) Number() uint64 { return uint64(f.num) }

func (f *Field[T]) Scan(r Reader) error { f.num.Scan(r); return f.Span.Scan(r) }
func (f Field[T]) Put(w Writer) error   { f.num.Put(w); return f.Span.Put(w) }
func (f Field[T]) Size() uint64         { return f.num.Size() + f.Span.Size() }

func (f Field[T]) PutText(w Writer) error {
	fmt.Fprintf(w, "<field n=%d, size=%d>\n", f.num, f.Span.Size())
	PutText(w, f.Elem())
	fmt.Fprintln(w, "</field>")
	return nil
}
