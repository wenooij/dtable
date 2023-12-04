package dtwire

import "io"

type Seq[T Scalar] []T

func (a Seq[T]) Put(w Writer) error {
	for _, e := range a {
		if err := any(e).(Putter).Put(w); err != nil {
			return err
		}
	}
	return nil
}

func (a *Seq[T]) Scan(r Reader) error {
	for b := make(Seq[T], 0, 8); ; {
		var t T
		if err := any(&t).(Scanner).Scan(r); err != nil {
			if err == io.EOF {
				*a = b
				return nil
			}
			return err
		}
		b = append(b, t)
	}
}
