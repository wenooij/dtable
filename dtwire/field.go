package dtwire

type Field[T Putter] struct {
	Number Uint64
	Value  Span[T]
}

func (f *Field[T]) Scan(r Reader) error {
	f.Number.Scan(r)
	return f.Value.Scan(r)
}

func (f Field[T]) Put(w Writer) error {
	f.Number.Put(w)
	return f.Value.Put(w)
}
