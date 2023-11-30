package dtable

type Collection interface {
	FirstTable() Table
	EachTable(func(Table) bool)
	PrefixTable(prefix []byte) Table
	LastTable() Table
}

type Table interface {
	FirstBlock() Block
	Block(i int) Block
	LastBlock() Block
	Prefix(prefix []byte) Block
}

type Block interface {
	PrevBlock() Block
	ScanRecords(func(v any) bool)
	NextBlock() Block
}

type Emitter interface {
	Emit(record any)
}

func ScanTable(t Table, fn func(any) bool) {
	var stop bool
	for b := t.FirstBlock(); b != nil && !stop; b = b.NextBlock() {
		b.ScanRecords(func(v any) bool {
			stop = !fn(v)
			return !stop
		})
	}
}

func Filter(e Emitter, t Table, fn func(any) bool) {
	ScanTable(t, func(v any) bool {
		if fn(v) {
			e.Emit(v)
		}
		return true
	})
}
