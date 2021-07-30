package types

type ErrorCode uint8

const (
	NO_ERROR ErrorCode = iota
	LOOP_UNCLOSED
	LOOP_UNOPNED
	MEM_OUT_OF_RANGE
	MEM_ROLL_OVER
)