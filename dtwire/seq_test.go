package dtwire

import (
	"bytes"
	"strings"
	"testing"
)

func TestFieldSeq(t *testing.T) {
	s := Seq[Span]{}
	s.Scan(strings.NewReader("\x09\x00\x01\x00\x01\x01\x01\x02\x01\x02"))
	for _, e := range s {
		ss := Tup3[Field, Field, Field]{}
		ss.Scan(bytes.NewReader(e))
		t.Log(ss)
	}
}
