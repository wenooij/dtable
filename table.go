package dtable

type Table interface {
	FirstBlock() Block
	Block(i int) Block
	LastBlock() Block
	Prefix(prefix []byte) Block
}

type Block interface {
	PrevBlock() Block
	ScanRecords(func(Consumer) bool)
	NextBlock() Block
}

type Consumer interface {
	Peek() (any, bool)
	Advance()
}

type Transaction interface {
	Revert()
	Commit()
}

type Emitter interface {
	Transaction
	Emit(record any)
}

type ConsumeEmitter interface {
	Consumer
	Emitter
}

func ScanTable(t Table, fn func(Consumer) bool) {
	var stop bool
	for b := t.FirstBlock(); b != nil && !stop; b = b.NextBlock() {
		b.ScanRecords(func(c Consumer) bool {
			stop = !fn(c)
			return !stop
		})
	}
}

func ProcessTable(e Emitter, t Table, fn func(ConsumeEmitter)) {
	ScanTable(t, func(c Consumer) bool {
		fn(struct {
			Consumer
			Emitter
		}{c, e})
		return true
	})
}
