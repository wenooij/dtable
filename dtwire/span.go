package dtwire

type Span[T Putter] struct {
	N     Uint64
	Value T
}

func SpanOf[T Putter](v T) Span[T] {
	return Span[T]{N: Uint64(v.Size()), Value: v}
}

func (x *Span[T]) Scan(r Reader) error {
	x.N.Scan(r)
	return any(&x.Value).(Scanner).Scan(r)
}

func (x Span[T]) Put(w Writer) error {
	x.N.Put(w)
	return x.Value.Put(w)
}

func (x Span[T]) Size() uint64 { return x.N.Size() + x.Value.Size() }
