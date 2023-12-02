package inmemory

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wenooij/dtable"
)

func TestFilter(t *testing.T) {
	b0 := &Table{}
	b0.Insert([]byte("a"))
	b0.Insert([]byte("ab"))
	b0.Insert([]byte("abc"))
	b0.Insert([]byte("bc"))
	b0.Insert([]byte("c"))

	b1 := &Table{}

	dtable.ProcessTable(b1, b0, func(e dtable.ConsumeEmitter) {
		defer e.Commit()
		defer e.Advance()
		v, _ := e.Peek()
		if bytes.HasPrefix(v.([]byte), []byte{'a'}) {
			e.Emit(v)
		}
	})

	dtable.ScanTable(b1, func(c dtable.Consumer) bool {
		defer c.Advance()
		v, _ := c.Peek()
		fmt.Println(string(v.([]byte)))
		return true
	})
}

func TestMap(t *testing.T) {
	b0 := &Table{}
	b0.Insert(uint64(1))
	b0.Insert(uint64(2))
	b0.Insert(uint64(3))
	b0.Insert(uint64(4))
	b0.Insert(uint64(5))

	b1 := &Table{}

	dtable.ProcessTable(b1, b0, func(e dtable.ConsumeEmitter) {
		defer e.Commit()
		defer e.Advance()
		v, _ := e.Peek()
		e.Emit(v.(uint64) * v.(uint64))
	})

	dtable.ScanTable(b1, func(c dtable.Consumer) bool {
		defer c.Advance()
		v, _ := c.Peek()
		fmt.Println(v.(uint64))
		return true
	})
}
