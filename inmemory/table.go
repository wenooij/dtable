package inmemory

import (
	"github.com/wenooij/dtable"
)

type TableBuilder struct {
	records []any
}

func (b *TableBuilder) Emit(v any) { b.Insert(v) }

func (b *TableBuilder) Insert(v any) {
	b.records = append(b.records, v)
}

func (b *TableBuilder) Build() *Table {
	return &Table{records: b.records}
}

type Table struct {
	records []any
}

// dtable.Table methods.

func (t *Table) FirstBlock() dtable.Block { return t }
func (t *Table) Block(i int) dtable.Block {
	if i == 0 {
		return t
	}
	return nil
}
func (t *Table) LastBlock() dtable.Block           { return t }
func (t *Table) Prefix(prefix []byte) dtable.Block { return t }

// dtable.Block methods.

func (*Table) PrevBlock() dtable.Block { return nil }
func (t *Table) ScanRecords(fn func(v any) bool) {
	for _, v := range t.records {
		if !fn(v) {
			break
		}
	}
}
func (*Table) NextBlock() dtable.Block { return nil }
