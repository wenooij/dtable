package dtwire

type Field[T Putter] struct {
	Number Uint64
	Value  Span[T]
}

func FieldOf[T Putter](n uint64, v T) Field[T] {
	return Field[T]{Number: Uint64(n), Value: SpanOf(v)}
}

func (f *Field[T]) Scan(r Reader) error { f.Number.Scan(r); return f.Value.Scan(r) }
func (f Field[T]) Put(w Writer) error   { f.Number.Put(w); return f.Value.Put(w) }
func (f Field[T]) Size() uint64         { return f.Number.Size() + f.Value.Size() }
