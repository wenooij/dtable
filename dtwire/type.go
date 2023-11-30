package dtwire

type typeNum byte

const (
	varIntType typeNum = iota
	fixed32Type
	fixed64Type
	lengthType
)
