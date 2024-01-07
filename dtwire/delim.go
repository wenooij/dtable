package dtwire

import (
	"fmt"
	"io"
	"math"
)

type Bytes raw

func (x *Bytes) Scan(r Reader) error {
	var n Uint64
	if err := n.Scan(r); err != nil {
		return err
	}
	if n > math.MaxInt64 {
		return ErrScan
	}
	return (*raw)(x).Scan(limitReader(r, int64(n)))
}

func (x Bytes) Put(w Writer) error {
	ux := Uint64(len(x))
	ux.Put(w)
	_, err := w.Write(x)
	return err
}

func (x Bytes) Size() uint64 { return Uint64(len(x)).Size() + raw(x).Size() }

func (x Bytes) PutText(w Writer) error {
	fmt.Fprintf(w, "<bytes size=%d x=%x />\n", len(x), x)
	return nil
}

type String Bytes

func (x *String) Scan(r Reader) error { return (*Bytes)(x).Scan(r) }
func (x String) Put(w Writer) error   { return Bytes(x).Put(w) }
func (x String) Size() uint64         { return Bytes(x).Size() }
func (x String) String() string       { return string(x) }
func (x String) PutText(w Writer) error {
	fmt.Fprintf(w, "<string size=%d x=%q />\n", len(x), string(x))
	return nil
}

type raw []byte

func (x *raw) Scan(r Reader) error {
	b := *x
	// Modified version of io.ReadAll with smaller read sizes.
	if cap(b) == 0 {
		b = make([]byte, 0, 8)
	}
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
	}

	// Store the bytes and trim extra capacity.
	*x = raw(b[0:len(b):len(b)])
	return nil
}

func (x raw) Put(w Writer) error {
	_, err := w.Write(x)
	return err
}

func (x raw) Size() uint64 { return uint64(len(x)) }

type limitedReader struct {
	r Reader
	n int64
}

func limitReader(r Reader, n int64) *limitedReader {
	return &limitedReader{r, n}
}

func (l *limitedReader) ReadByte() (b byte, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	b, err = l.r.ReadByte()
	if err != nil {
		return 0, err
	}
	l.n--
	return b, nil
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}
