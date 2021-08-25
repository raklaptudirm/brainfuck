/*
Package types provides the types required by the parser.
*/
package types

// Instruction type represents a single brainfuck instuction
type Instruction uint8

// LoopIndexes type represents the index of a loop,
// utilized to improve runtime speeds.
type LoopIndexes int

// Instructions representing brainfuck commands.
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
