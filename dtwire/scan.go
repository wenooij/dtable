package dtwire

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"golang.org/x/exp/constraints"
)

var ErrScan = errors.New("scan error")

func ScanVarint(r *bufio.Reader) (int64, error) {
	ux, err := ScanUvarint(r)
	if err != nil {
		return 0, err
	}
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x, nil
}

func ScanUvarint(r *bufio.Reader) (uint64, error) {
	var x uint64
	var s uint
	for i := 0; i < binary.MaxVarintLen64; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b < 0x80 {
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
	return 0, io.ErrShortBuffer
}

func ScanFixed32(r *bufio.Reader) (uint32, error) {
	b0, _ := r.ReadByte()
	b1, _ := r.ReadByte()
	b2, _ := r.ReadByte()
	b3, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint32(b0) | uint32(b1)<<8 | uint32(b2)<<16 | uint32(b3)<<24, nil
}

func ScanFixed64(r *bufio.Reader) (uint64, error) {
	b0, _ := r.ReadByte()
	b1, _ := r.ReadByte()
	b2, _ := r.ReadByte()
	b3, _ := r.ReadByte()
	b4, _ := r.ReadByte()
	b5, _ := r.ReadByte()
	b6, _ := r.ReadByte()
	b7, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint64(b0) | uint64(b1)<<8 | uint64(b2)<<16 | uint64(b3)<<24 | uint64(b4)<<32 | uint64(b5)<<40 | uint64(b6)<<48 | uint64(b7)<<56, nil
}

func ScanBool(r *bufio.Reader) (bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return false, err
	}
	switch b {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, ErrScan
	}
}

func ScanUnsigned[T constraints.Unsigned](r *bufio.Reader) (T, error) {
	v, err := ScanUvarint(r)
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

func ScanSigned[T constraints.Signed](r *bufio.Reader) (T, error) {
	v, err := ScanVarint(r)
	if err != nil {
		return 0, err
	}
	return T(v), nil
}

func ScanBytes(r *bufio.Reader) ([]byte, error) {
	n, err := ScanUvarint(r)
	if err != nil {
		return nil, err
	}
	return r.Peek(int(n))
}

func ScanField(w *bufio.Reader) (uint32, typeNum, error) {
	return 0, 0, fmt.Errorf("not implemented")
}
