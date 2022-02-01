/*
Package types provides types for error codes.
*/
package types

// ErrorCode is the datatype for brainfuck
// parse error codes.
type ErrorCode uint8

// Error codes returned by the parser on
// a parsing error.
const (
	NO_ERROR ErrorCode = iota
	LOOP_UNCLOSED
	LOOP_UNOPNED
	MEM_OUT_OF_RANGE
	MEM_ROLL_OVER
)
