package inmemory

import (
	"github.com/wenooij/dtable"
)

type Table struct {
	records []any

	tx []any

	p int
}

func (t *Table) Reset()         { t.records = t.records[:0]; t.ResetConsumer() }
func (t *Table) ResetConsumer() { t.p = 0 }
func (t *Table) ResetEmitter()  { t.tx = t.tx[:0] }
func (t *Table) Revert()        { t.ResetEmitter() }
func (t *Table) Commit()        { t.records = append(t.records, t.tx...); t.ResetEmitter() }

func (t *Table) Emit(v any)   { t.tx = append(t.tx, v) }
func (t *Table) Insert(v any) { t.records = append(t.records, v) }

func (t *Table) Peek() (any, bool) {
	if t.p < len(t.records) {
		return t.records[t.p], true
	}
	return nil, false
}

func (t *Table) Advance() {
	if t.p < len(t.records) {
		t.p++
	}
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
func (t *Table) ScanRecords(fn func(e dtable.Consumer) bool) {
	for t.p < len(t.records) {
		if !fn(t) {
			break
		}
	}
}
func (*Table) NextBlock() dtable.Block { return nil }
