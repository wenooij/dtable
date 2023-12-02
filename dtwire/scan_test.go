package dtwire

import (
	"bufio"
	"bytes"
	"testing"
)

func TestScanPair(t *testing.T) {
	var f Field[any]
	err := ScanField(bufio.NewReader(bytes.NewReader([]byte("\x10\x00\x00"))), &f)
	t.Log(f, err)
}
