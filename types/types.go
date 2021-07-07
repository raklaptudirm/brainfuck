package types

const TAPE_LENGTH uint16 = 30000

type Instruction uint8

const (
	GO_LEFT = iota
	GO_RIGHT
	INCREMENT
	DECREMENT
	INPUT
	OUTPUT
	LOOP_START
	LOOP_END
)
