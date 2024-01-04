package dtwire

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
