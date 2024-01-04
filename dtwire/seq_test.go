package dtwire

import (
	"strings"
	"testing"
)

func TestFieldSeq(t *testing.T) {
	s := Seq[Field[Uint64]]{}
	s.Scan(strings.NewReader("\x09\x00\x01\x00\x01\x01\x01\x02\x01\x02"))
	t.Log(s)
	for _, e := range s {
		t.Log(e)
	}
}
