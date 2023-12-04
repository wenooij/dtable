package dtwire

type Tup[T Scalar] Seq[T]

func (a Tup[T]) Put(w Writer) error { return Seq[T](a).Put(w) }

func (a Tup[T]) Scan(r Reader) error {
	for i := range a {
		if err := any(&a[i]).(Scanner).Scan(r); err != nil {
			return err
		}
	}
	return nil
}
