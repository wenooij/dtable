package dtwire

import (
	"errors"
	"io"
)

var (
	ErrType = errors.New("type error")
	ErrScan = errors.New("scan error")
)

type Writer interface {
	io.Writer
	io.ByteWriter
}

type Reader interface {
	io.Reader
	io.ByteReader
}

type Scanner interface {
	Scan(Reader) error
}

type Putter interface {
	Put(Writer) error
}

type ScanPutter interface {
	Scanner
	Putter
}

type Varint interface {
	Bool | Int64 | Uint64
}

type Fixed interface {
	Byte | Fixed32 | Fixed64 | Float32 | Float64
}

type Delimited interface {
	Bytes | String | Field | Span
}

type Scalar interface {
	Varint | Fixed | Delimited
}
