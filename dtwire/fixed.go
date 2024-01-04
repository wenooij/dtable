package dtwire

import "math"

type Byte byte

func (b *Byte) Scan(r Reader) error {
	bb, err := r.ReadByte()
	if err != nil {
		return err
	}
	*b = Byte(bb)
	return nil
}

func (b Byte) Put(w Writer) error {
	return w.WriteByte(byte(b))
}

type Bool bool

func (b *Bool) Scan(r Reader) error {
	bb, err := r.ReadByte()
	if err != nil {
		return err
	}
	switch bb {
	case 0:
		*b = false
		return nil
	case 1:
		*b = true
		return nil
	default:
		return ErrScan
	}
}

func (x Bool) Put(w Writer) error {
	var b byte
	if x {
		b = 1
	}
	return w.WriteByte(b)
}

type Fixed32 uint32

func (i *Fixed32) Scan(r Reader) error {
	b0, _ := r.ReadByte()
	b1, _ := r.ReadByte()
	b2, _ := r.ReadByte()
	b3, err := r.ReadByte()
	if err != nil {
		return err
	}
	*i = Fixed32(uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24)
	return nil
}

func (x Fixed32) Put(w Writer) error {
	w.WriteByte(byte(x))
	w.WriteByte(byte(x >> 8))
	w.WriteByte(byte(x >> 16))
	return w.WriteByte(byte(x >> 24))
}

type Fixed64 uint64

func (i *Fixed64) Scan(r Reader) error {
	b0, _ := r.ReadByte()
	b1, _ := r.ReadByte()
	b2, _ := r.ReadByte()
	b3, _ := r.ReadByte()
	b4, _ := r.ReadByte()
	b5, _ := r.ReadByte()
	b6, _ := r.ReadByte()
	b7, err := r.ReadByte()
	if err != nil {
		return err
	}
	*i = Fixed64(uint64(b0) | uint64(b1)<<8 | uint64(b2)<<16 | uint64(b3)<<24 | uint64(b4)<<32 | uint64(b5)<<40 | uint64(b6)<<48 | uint64(b7)<<56)
	return nil
}

func (x Fixed64) Put(w Writer) error {
	w.WriteByte(byte(x))
	w.WriteByte(byte(x >> 8))
	w.WriteByte(byte(x >> 16))
	w.WriteByte(byte(x >> 24))
	w.WriteByte(byte(x >> 32))
	w.WriteByte(byte(x >> 40))
	w.WriteByte(byte(x >> 48))
	return w.WriteByte(byte(x >> 56))
}

type Float32 float32

func (x *Float32) Scan(r Reader) error {
	var i Fixed32
	if err := i.Scan(r); err != nil {
		return err
	}
	*x = Float32(math.Float32frombits(uint32(i)))
	return nil
}

func (x Float32) Put(w Writer) error {
	i := Fixed32(math.Float32bits(float32(x)))
	return i.Put(w)
}

type Float64 float64

func (x *Float64) Scan(r Reader) error {
	var i Fixed64
	if err := i.Scan(r); err != nil {
		return err
	}
	*x = Float64(math.Float64frombits(uint64(i)))
	return nil
}

func (x Float64) Put(w Writer) error {
	i := Fixed64(math.Float64bits(float64(x)))
	return i.Put(w)
}
