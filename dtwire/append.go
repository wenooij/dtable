package dtwire

import (
	"bufio"
	"math"

	"golang.org/x/exp/constraints"
)

func AppendVarint(w *bufio.Writer, x int64) {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	AppendUvarint(w, ux)
}

func AppendUvarint(w *bufio.Writer, x uint64) {
	for x >= 0x80 {
		w.WriteByte(byte(x) | 0x80)
		x >>= 7
	}
	w.WriteByte(byte(x))
}

func AppendFixed32(w *bufio.Writer, x uint32) {
	w.WriteByte(byte(x))
	w.WriteByte(byte(x >> 8))
	w.WriteByte(byte(x >> 16))
	w.WriteByte(byte(x >> 24))
}

func AppendFixed64(w *bufio.Writer, x uint64) {
	w.WriteByte(byte(x))
	w.WriteByte(byte(x >> 8))
	w.WriteByte(byte(x >> 16))
	w.WriteByte(byte(x >> 24))
	w.WriteByte(byte(x >> 32))
	w.WriteByte(byte(x >> 40))
	w.WriteByte(byte(x >> 48))
	w.WriteByte(byte(x >> 56))
}

func AppendBool(w *bufio.Writer, x bool) {
	var b byte
	if x {
		b = 1
	}
	w.WriteByte(b)
}

func AppendUnsigned[T constraints.Unsigned](w *bufio.Writer, x T) { AppendUvarint(w, uint64(x)) }
func AppendSigned[T constraints.Signed](w *bufio.Writer, x T)     { AppendVarint(w, int64(x)) }

func AppendFloat32(w *bufio.Writer, x float32) { AppendFixed32(w, math.Float32bits(x)) }
func AppendFloat64(w *bufio.Writer, x float64) { AppendFixed64(w, math.Float64bits(x)) }

func AppendBytes(w *bufio.Writer, x []byte) {
	AppendUvarint(w, uint64(len(x)))
	w.Write(x)
}

func AppendField(w *bufio.Writer, fieldNumber uint32, typeNum typeNum) {
	AppendUvarint(w, uint64(fieldNumber)<<3|uint64(typeNum))
}
