package dtwire

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"math"
)

var ErrScan = errors.New("scan error")

func scanBool(r *bufio.Reader, b *bool) error {
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

func scanVarint(r *bufio.Reader, i *int64) error {
	var ux uint64
	if err := scanUvarint(r, &ux); err != nil {
		return err
	}
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	*i = x
	return nil
}

func scanUvarint(r *bufio.Reader, e *uint64) error {
	var x uint64
	var s uint
	for i := 0; i < binary.MaxVarintLen64; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return err
		}
		if b < 0x80 {
			*e = x | uint64(b)<<s
			return nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
	return io.ErrShortBuffer
}

type Fixed32 uint32

func scanFixed32(r *bufio.Reader, i *Fixed32) error {
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

type Fixed64 uint64

func scanFixed64(r *bufio.Reader, i *Fixed64) error {
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

func scanLength[T any](r *bufio.Reader, e *T) error {
	var n uint64
	if err := scanUvarint(r, &n); err != nil {
		return err
	}
	bb, err := r.Peek(int(n))
	if err != nil {
		return err
	}
	switch e := any(e).(type) {
	case *[]byte:
		*e = bb
		return nil
	case *string:
		*e = string(bb)
		return nil
	default:
		return ErrType
	}
}

type Field[T any] struct {
	FieldNumber int32
	Value       T
}

func ScanField[T any](r *bufio.Reader, field *Field[T]) error {
	var ux uint64
	if err := scanUvarint(r, &ux); err != nil {
		return err
	}
	typeNum := ux & 7
	ux >>= 3
	if ux > math.MaxInt32 {
		return ErrScan
	}
	field.FieldNumber = int32(ux)
	switch typeNum {
	case 0: // varInt
		switch v := any(&field.Value).(type) {
		case *int64:
			return scanVarint(r, v)
		case *uint64:
			return scanUvarint(r, v)
		case *any:
			var ux uint64
			if err := scanUvarint(r, &ux); err != nil {
				return err
			}
			*v = ux
			return nil
		default:
			return ErrType
		}
	case 1: // fixed32
		switch v := any(&field.Value).(type) {
		case *Fixed32:
			return scanFixed32(r, v)
		default:
			return ErrType
		}
	case 2: // fixed64
		switch v := any(&field.Value).(type) {
		case *Fixed64:
			return scanFixed64(r, v)
		default:
			return ErrType
		}
	case 3: // length
		switch v := any(&field.Value).(type) {
		case *[]byte:
			return scanLength(r, v)
		case *string:
			return scanLength(r, v)
		default:
			return ErrType
		}
	default:
		return ErrScan
	}
}

var ErrType = errors.New("type error")

func Scan[T any](r *bufio.Reader, e *T) error {
	switch e := any(e).(type) {
	case *bool:
		return scanBool(r, e)
	case *int64:
		return scanVarint(r, e)
	case *uint64:
		return scanUvarint(r, e)
	case *Fixed32:
		return scanFixed32(r, e)
	case *Fixed64:
		return scanFixed64(r, e)
	case *float32:
		var i Fixed32
		if err := scanFixed32(r, &i); err != nil {
			return err
		}
		*e = math.Float32frombits(uint32(i))
		return nil
	case *float64:
		var i Fixed64
		if err := scanFixed64(r, &i); err != nil {
			return err
		}
		*e = math.Float64frombits(uint64(i))
		return nil
	case *[]byte:
		return scanLength(r, e)
	case *string:
		return scanLength(r, e)
	case *Field[T]:
		return ScanField[T](r, e)
	default:
		return ErrType
	}
}
