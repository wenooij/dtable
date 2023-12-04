package dtwire

type Tup2[T0, T1 Scalar] struct {
	T0 T0
	T1 T1
}

func (a Tup2[T0, T1]) Put(w Writer) error {
	any(a.T0).(Putter).Put(w)
	return any(a.T1).(Putter).Put(w)
}

func (a *Tup2[T0, T1]) Scan(r Reader) error {
	any(&a.T0).(Scanner).Scan(r)
	return any(&a.T1).(Scanner).Scan(r)
}

type Tup3[T0, T1, T2 Scalar] struct {
	T0 T0
	T1 T1
	T2 T2
}

func (a Tup3[T0, T1, T2]) Put(w Writer) error {
	any(&a.T0).(Putter).Put(w)
	any(&a.T1).(Putter).Put(w)
	return any(&a.T2).(Putter).Put(w)
}

func (a *Tup3[T0, T1, T2]) Scan(r Reader) error {
	any(&a.T0).(Scanner).Scan(r)
	any(&a.T1).(Scanner).Scan(r)
	return any(&a.T2).(Scanner).Scan(r)
}
