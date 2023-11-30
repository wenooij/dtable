package inmemory

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wenooij/dtable"
)

func TestFilter(t *testing.T) {
	b0 := &TableBuilder{}
	b0.Insert([]byte("a"))
	b0.Insert([]byte("ab"))
	b0.Insert([]byte("abc"))
	b0.Insert([]byte("bc"))
	b0.Insert([]byte("c"))

	b1 := &TableBuilder{}

	dtable.Filter(b1, b0.Build(), func(v any) bool {
		return bytes.HasPrefix(v.([]byte), []byte{'a'})
	})

	dtable.ScanTable(b1.Build(), func(v any) bool {
		fmt.Println(string(v.([]byte)))
		return true
	})
}
