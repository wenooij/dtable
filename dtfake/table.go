package dtfake

import (
	"bytes"
	"io"

	"github.com/wenooij/dtable/dtwire"
)

type tableScanner struct {
	*Table
	b bytes.Buffer
	p int
}

func (t *tableScanner) Scan(v dtwire.Scanner) (err error) {
	if len(t.records) <= t.p {
		return io.EOF
	}
	t.b.Reset()
	if err := t.records[t.p].Put(&t.b); err != nil {
		return err
	}
	defer func() {
		if err == nil {
			t.p++
		}
	}()
	return v.Scan(&t.b)
}

// Table represents a fake Table in which dtwire primitives can be directly used.
type Table struct {
	records []dtwire.Putter
	*tableScanner
}

func (t *Table) Reset()              { t.records = t.records[:0]; t.ResetScan() }
func (t *Table) ResetScan()          { t.tableScanner = nil }
func (t *Table) Put(v dtwire.Putter) { t.records = append(t.records, v) }
func (t *Table) Scan(v dtwire.Scanner) error {
	if t.tableScanner == nil {
		t.tableScanner = &tableScanner{Table: t}
	}
	return t.tableScanner.Scan(v)
}
