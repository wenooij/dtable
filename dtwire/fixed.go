package dtwire

import (
	"fmt"
	"math"
)

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

func (Byte) Size() uint64 { return 1 }

func (b Byte) PutText(w Writer) error { fmt.Fprintf(w, "<byte x=%x />\n", b); return nil }

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

func (x Bool) Size() uint64 { return 1 }

func (b Bool) PutText(w Writer) error { fmt.Fprintf(w, "<%v />\n", b); return nil }

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

func (x Fixed32) Size() uint64 { return 4 }

func (x Fixed32) PutText(w Writer) error { fmt.Fprintf(w, "<fixed32 x=%d />\n", x); return nil }

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

func (x Fixed64) Size() uint64 { return 8 }

func (x Fixed64) PutText(w Writer) error { fmt.Fprintf(w, "<fixed64 x=%d />\n", x); return nil }

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

func (x Float32) Size() uint64 { return 4 }

func (x Float32) PutText(w Writer) error { fmt.Fprintf(w, "<float32 x=%f />\n", x); return nil }

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

func (x Float64) Size() uint64 { return 8 }

func (x Float64) PutText(w Writer) error { fmt.Fprintf(w, "<float64 x=%f />\n", x); return nil }
