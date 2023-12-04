package dtwire

type Field struct {
	Number Uint64
	Value  Span
}

func (f *Field) Scan(r Reader) error {
	f.Number.Scan(r)
	return f.Value.Scan(r)
}

func (f Field) Put(w Writer) error {
	f.Number.Put(w)
	return f.Value.Put(w)
}
