// brainfuck
// https://github.com/raklaptudirm/brainfuck
// Copyright (c) 2021 Rak Laptudirm.
// Licensed under the MIT license.

// Package parser provides parsing functions for a brainfuck source.
//
// The provided method and types is used for parsing a brainfuck source
// string into an Instruction slice.
//
package parser

import (
	"errors"
	"fmt"

	. "github.com/raklaptudirm/brainfuck/pkg/errors/types"
)

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

type Bytecode struct {
	Instructions []Instruction
	Indexes      []LoopIndexes
}

// Parse parses the given brainfuck source,
// and returns the instruction codes, loop indexes,
//and errors(if any) found in the source string.
func Parse(code string) (Bytecode, error, ErrorCode) {
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

	ret := Bytecode{Instructions: bytecode, Indexes: indexes}

	return ret, parseError, errorCode
}
