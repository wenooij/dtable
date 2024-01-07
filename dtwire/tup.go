package dtwire

import "fmt"

type Tup2[T0, T1 Putter] struct {
	T0 T0
	T1 T1
}

func (a Tup2[T0, T1]) Put(w Writer) error {
	a.T0.Put(w)
	return a.T1.Put(w)
}

func (a *Tup2[T0, T1]) Scan(r Reader) error {
	any(&a.T0).(Scanner).Scan(r)
	return any(&a.T1).(Scanner).Scan(r)
}

func (a Tup2[T0, T1]) Size() uint64 { return a.T0.Size() + a.T1.Size() }

func (a Tup2[T0, T1]) PutText(w Writer) error {
	fmt.Fprintln(w, "<tup2>")
	PutText(w, a.T0)
	PutText(w, a.T1)
	fmt.Fprintln(w, "</tup2>")
	return nil
}

type Tup3[T0, T1, T2 Putter] struct {
	T0 T0
	T1 T1
	T2 T2
}

func (a Tup3[T0, T1, T2]) Put(w Writer) error {
	a.T0.Put(w)
	a.T1.Put(w)
	return a.T2.Put(w)
}

func (a *Tup3[T0, T1, T2]) Scan(r Reader) error {
	any(&a.T0).(Scanner).Scan(r)
	any(&a.T1).(Scanner).Scan(r)
	return any(&a.T2).(Scanner).Scan(r)
}

func (a Tup3[T0, T1, T2]) Size() uint64 { return a.T0.Size() + a.T1.Size() + a.T2.Size() }

func (a Tup3[T0, T1, T2]) PutText(w Writer) error {
	fmt.Fprintln(w, "<tup3>")
	PutText(w, a.T0)
	PutText(w, a.T1)
	PutText(w, a.T2)
	fmt.Fprintln(w, "</tup3>")
	return nil
}
