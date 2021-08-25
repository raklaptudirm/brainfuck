/*
Package parser provides parsing functions for a brainfuck source.
*/
package parser

import (
	"errors"
	"fmt"

	. "github.com/raklaptudirm/brainfuck/errors/types"
	. "github.com/raklaptudirm/brainfuck/parser/types"
)

// Parse parses the given brainfuck source,
// and returns the instruction codes, loop indexes,
//and errors(if any) found in the source string.
func Parse(code string) ([]Instruction, error, []LoopIndexes, ErrorCode) {
	var parseError error = nil
	var errorCode ErrorCode = NO_ERROR

	bytecode := []Instruction{}
	indexes := []LoopIndexes{}
	loops := []LoopIndexes{}

	length := len(code)
	lines := 1
	column := 0

	last := func(item []LoopIndexes) int {
		return len(item) - 1
	}

	elements := func() int {
		return len(bytecode) - 1
	}

	for i := 0; i < length; i += 1 {
		column += 1

		indexes = append(indexes, 0)

		switch string(code[i]) {
		case "<":
			bytecode = append(bytecode, GO_LEFT)
		case ">":
			bytecode = append(bytecode, GO_RIGHT)
		case ".":
			bytecode = append(bytecode, OUTPUT)
		case ",":
			bytecode = append(bytecode, INPUT)
		case "+":
			bytecode = append(bytecode, INCREMENT)
		case "-":
			bytecode = append(bytecode, DECREMENT)
		case "[":
			bytecode = append(bytecode, LOOP_START)

			// The loops start's index is the number of elements in bytecode - 1
			// since the command has already been pushed into the array.
			// The index is converted to parser/types.LoopIndexes,
			// and appended to the loops array.
			loops = append(loops, LoopIndexes(elements()))
		case "]":
			if len(loops) == 0 {
				parseError = errors.New(fmt.Sprintf("error %v:%v : Illeagal \"]\".", lines, column))
				errorCode = LOOP_UNOPNED
			} else {
				bytecode = append(bytecode, LOOP_END)

				// The index of the starting brace of this loop,
				// will be the last element of loops
				loopStart := loops[last(loops)]

				// Switch the indexes of the starting and ending brace,
				// and add them to indexes.
				indexes[loopStart] = LoopIndexes(elements())
				indexes[elements()] = loopStart

				loops = loops[:last(loops)]
			}
		case "\n":
			lines += 1
			column = 0
			indexes = indexes[:last(indexes)]
		default:
			indexes = indexes[:last(indexes)]
		}
	}

	if len(loops) != 0 {
		parseError = errors.New(fmt.Sprintf("error: %v unclosed \"[\".", len(loops)))
		errorCode = LOOP_UNCLOSED
	}

	return bytecode, parseError, indexes, errorCode
}
