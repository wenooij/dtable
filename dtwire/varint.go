package dtwire

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Uint64 uint64

func (e *Uint64) Scan(r Reader) error {
	var x Uint64
	var s uint
	for i := 0; i < binary.MaxVarintLen64; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return err
		}
		if b < 0x80 {
			*e = x | Uint64(b)<<s
			return nil
		}
		x |= Uint64(b&0x7f) << s
		s += 7
	}
	return io.ErrShortBuffer
}

func (x Uint64) Put(w Writer) error {
	for i := 0; x >= 0x80; i++ {
		w.WriteByte(byte(x) | 0x80)
		x >>= 7
	}
	return w.WriteByte(byte(x))
}

func (x Uint64) Size() uint64 {
	n := uint64(1)
	for i := 0; x >= 0x80; i++ {
		n++
		x >>= 7
	}
	return n
}

func (x Uint64) PutText(w Writer) error {
	fmt.Fprintf(w, "<uint x=%d />\n", x)
	return nil
}

type Int64 int64

func (i *Int64) Scan(r Reader) error {
	var ux Uint64
	if err := ux.Scan(r); err != nil {
		return err
	}
	x := Int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	*i = x
	return nil
}

func (x Int64) Put(w Writer) error {
	ux := Uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return ux.Put(w)
}

func (x Int64) Size() uint64 {
	ux := Uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return ux.Size()
}

func (x Int64) PutText(w Writer) error {
	fmt.Fprintf(w, "<int x=%d />\n", x)
	return nil
}
