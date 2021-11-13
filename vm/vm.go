// brainfuck
// https://github.com/raklaptudirm/brainfuck
// Copyright (c) 2021 Rak Laptudirm.
// Licensed under the MIT license.

// Package vm provides methods and structures to run
// parsed brainfuck code. It uses the Instruction
// slice received from parsing the source and runs
// each instruction individually.
//
package vm

import (
	"fmt"
	"io"

  "github.com/raklaptudirm/brainfuck/parser"
	. "github.com/raklaptudirm/brainfuck/vm/types"
)

// VM struct to store memory information during runtime.
type VM struct {
	DataPointer uint16            // Index of the currently active cell.
	Memory      [TAPE_LENGTH]byte // The array of cells.
}

// Default represents the default configuration for a VM.
var Default VM = VM{DataPointer: 0, Memory: [TAPE_LENGTH]byte{}}

func (vm *VM) execute(out io.Writer, instruction parser.Instruction) {
	switch instruction {
	case parser.GO_LEFT:
		// If statement for roll over if pointer is at 0.
		if vm.DataPointer == 0 {
			vm.DataPointer = TAPE_LENGTH - 1
		} else {
			vm.DataPointer -= 1
		}
	case parser.GO_RIGHT:
		// If statement for roll over if pointer is at TAPE_LENGTH.
		if vm.DataPointer == TAPE_LENGTH-1 {
			vm.DataPointer = 0
		} else {
			vm.DataPointer += 1
		}
	case parser.INPUT:
		_, _ = fmt.Scanf("%c", &vm.Memory[vm.DataPointer])
	case parser.OUTPUT:
		fmt.Fprint(out, string(vm.Memory[vm.DataPointer]))
	case parser.INCREMENT:
		vm.Memory[vm.DataPointer] += 1
	case parser.DECREMENT:
		vm.Memory[vm.DataPointer] -= 1
	}
}

// RunCode runs a parsed brainfuck source, consisting of
// the command array and the loop index array. The output
// is written to the provided io.Writer.
func (vm *VM) RunCode(out io.Writer, bytecode parser.Bytecode) {
	instructions := bytecode.Instructions
	indexes := bytecode.Indexes
	length := len(instructions)

	for i := 0; i < length; i += 1 {
		switch instructions[i] {
		case parser.LOOP_START:
			if vm.Memory[vm.DataPointer] == 0 {
				// Get end brace index from indexes array.
				i = int(indexes[i])
			}
		case parser.LOOP_END:
			if vm.Memory[vm.DataPointer] != 0 {
				// Get sart brace index from indexes array.
				i = int(indexes[i])
			}
		default:
			vm.execute(out, instructions[i])
		}
	}
}
