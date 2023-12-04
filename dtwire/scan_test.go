package dtwire

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newByte(v Byte) *Byte          { return &v }
func newBool(v Bool) *Bool          { return &v }
func newUint64(v Uint64) *Uint64    { return &v }
func newInt64(v Int64) *Int64       { return &v }
func newFixed32(v Fixed32) *Fixed32 { return &v }
func newFixed64(v Fixed64) *Fixed64 { return &v }
func newFloat32(v Float32) *Float32 { return &v }
func newFloat64(v Float64) *Float64 { return &v }
func newBytes(v Bytes) *Bytes       { return &v }
func newString(v String) *String    { return &v }

func newBufferedReader(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}

type scanTest struct {
	name        string
	inputReader Reader
	inputValue  ScanPutter
	wantValue   Putter
	wantErr     bool
}

func (tc scanTest) runTest(t *testing.T, scanPosfix string) {
	gotX := tc.inputValue
	err := gotX.Scan(tc.inputReader)
	if gotErr := err != nil; gotErr != tc.wantErr {
		t.Fatalf("Scan%s(%q): got err = %v, want err = %v", scanPosfix, tc.name, err, tc.wantErr)
	}
	if diff := cmp.Diff(tc.wantValue, gotX); diff != "" {
		t.Errorf("Scan%s(%q): got diff (-want +got):\n%v", scanPosfix, tc.name, diff)
	}
}

func TestScanScalar(t *testing.T) {
	for _, tc := range []scanTest{{
		name:        "scan byte",
		inputReader: newBufferedReader("\xff"),
		inputValue:  new(Byte),
		wantValue:   newByte('\xff'),
	}, {
		name:        "scan bool",
		inputReader: newBufferedReader("\x00"),
		inputValue:  new(Bool),
		wantValue:   newBool(false),
	}, {
		name:        "scan uvarint",
		inputReader: newBufferedReader("\x01"),
		inputValue:  new(Uint64),
		wantValue:   newUint64(1),
	}, {
		name:        "scan varint",
		inputReader: newBufferedReader("\x01"),
		inputValue:  new(Int64),
		wantValue:   newInt64(-1),
	}, {
		name:        "scan fixed32",
		inputReader: newBufferedReader("\x03\x00\x00\x00"),
		inputValue:  new(Fixed32),
		wantValue:   newFixed32(3),
	}, {
		name:        "scan fixed64",
		inputReader: newBufferedReader("\x03\x00\x00\x00\x00\x00\x00\x00"),
		inputValue:  new(Fixed64),
		wantValue:   newFixed64(3),
	}, {
		name:        "scan float32",
		inputReader: newBufferedReader("\xcd\xcc\xcc\x3d"),
		inputValue:  new(Float32),
		wantValue:   newFloat32(0.1),
	}, {
		name:        "scan float64",
		inputReader: newBufferedReader("\x9a\x99\x99\x99\x99\x99\xb9\x3f"),
		inputValue:  new(Float64),
		wantValue:   newFloat64(0.1),
	}, {
		name:        "scan bytes",
		inputReader: newBufferedReader("\x04\x00\x00\x00\x00"),
		inputValue:  new(Bytes),
		wantValue:   newBytes(Bytes("\x00\x00\x00\x00")),
	}, {
		name:        "scan string",
		inputReader: newBufferedReader("\x04\x00\x00\x00\x00"),
		inputValue:  new(String),
		wantValue:   newString(String("\x00\x00\x00\x00")),
	}} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.runTest(t, "Scalar")
		})
	}
}
