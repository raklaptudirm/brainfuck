package types

type Instruction uint8
type LoopIndexes int

const (
	GO_LEFT Instruction = iota
	GO_RIGHT
	INCREMENT
	DECREMENT
	INPUT
	OUTPUT
	LOOP_START
	LOOP_END
)
