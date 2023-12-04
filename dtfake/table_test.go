package dtfake

import (
	"io"
	"strings"
	"testing"

	"github.com/wenooij/dtable/dtwire"
)

func TestFilter(t *testing.T) {
	b0 := &Table{}
	b0.Put(dtwire.String("a"))
	b0.Put(dtwire.String("ab"))
	b0.Put(dtwire.String("abc"))
	b0.Put(dtwire.String("bc"))
	b0.Put(dtwire.String("c"))

	b1 := &Table{}

	for e := dtwire.String(""); ; {
		if err := b0.Scan(&e); err != nil {
			if err == io.EOF {
				break
			}
		}
		if strings.HasPrefix(string(e), "a") {
			b1.Put(e)
		}
	}

	for e := dtwire.String(""); ; {
		if err := b1.Scan(&e); err != nil {
			if err == io.EOF {
				break
			}
		}
		t.Log(e)
	}
}

func TestMap(t *testing.T) {
	b0 := &Table{}
	b0.Put(dtwire.Uint64(1))
	b0.Put(dtwire.Uint64(2))
	b0.Put(dtwire.Uint64(3))
	b0.Put(dtwire.Uint64(4))
	b0.Put(dtwire.Uint64(5))

	b1 := &Table{}

	for u := dtwire.Uint64(0); ; {
		if err := b0.Scan(&u); err != nil {
			if err == io.EOF {
				break
			}
		}
		b1.Put(u * u)
	}

	for u := dtwire.Uint64(0); ; {
		if err := b1.Scan(&u); err != nil {
			if err == io.EOF {
				break
			}
		}
		t.Log(u)
	}
}
