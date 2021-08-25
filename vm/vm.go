/*
Package vm provides methods and structures to run parsed brainfuck code.
*/
package vm

import (
	"fmt"
	. "io"
	"io/ioutil"

	. "github.com/raklaptudirm/brainfuck/errors"
	. "github.com/raklaptudirm/brainfuck/parser"
	. "github.com/raklaptudirm/brainfuck/parser/types"
	. "github.com/raklaptudirm/brainfuck/vm/types"
)

// VM struct to store memory information during runtime.
type VM struct {
	DataPointer uint16
	Memory      [TAPE_LENGTH]byte
}

// VMBase represents the default configuration for a VM
var VMBase VM = VM{DataPointer: 0, Memory: [TAPE_LENGTH]byte{}}

func (vm *VM) execute(out Writer, instruction Instruction) {
	switch instruction {
	case GO_LEFT:
		// If statement for roll over if pointer is at 0.
		if vm.DataPointer == 0 {
			vm.DataPointer = TAPE_LENGTH - 1
		} else {
			vm.DataPointer -= 1
		}
	case GO_RIGHT:
		// If statement for roll over if pointer is at TAPE_LENGTH
		if vm.DataPointer == TAPE_LENGTH-1 {
			vm.DataPointer = 0
		} else {
			vm.DataPointer += 1
		}
	case INPUT:
		_, _ = fmt.Scanf("%c", &vm.Memory[vm.DataPointer])
	case OUTPUT:
		fmt.Fprint(out, string(vm.Memory[vm.DataPointer]))
	case INCREMENT:
		vm.Memory[vm.DataPointer] += 1
	case DECREMENT:
		vm.Memory[vm.DataPointer] -= 1
	}
}

// RunCode runs a parsed brainfuck source, consisting of
// the command array and the loop index array. The output
// is written to the provided io.Writer
func (vm *VM) RunCode(out Writer, instructions []Instruction, indexes []LoopIndexes) {
	length := len(instructions)

	for i := 0; i < length; i += 1 {
		switch instructions[i] {
		case LOOP_START:
			if vm.Memory[vm.DataPointer] == 0 {
				// Get end brace index from indexes array.
				i = int(indexes[i])
			}
		case LOOP_END:
			if vm.Memory[vm.DataPointer] != 0 {
				// Get sart brace index from indexes array.
				i = int(indexes[i])
			}
		default:
			vm.execute(out, instructions[i])
		}
	}
}

// RunString runs a given brainfuck source string,
// by parsing it and then using VM.RunCode.
func (vm *VM) RunString(str string, out Writer) {
	instructions, parseError, indexes, _ := Parse(str)
	StrictCheck(parseError)

	vm.RunCode(out, instructions, indexes)
}

// RunFile runs a brainfuck source file, by reading it,
// parsing it, and then using VM.RunCode.
func (vm *VM) RunFile(fileName string, out Writer) {
	file, fileError := ioutil.ReadFile(fileName)
	StrictCheck(fileError)

	vm.RunString(string(file), out)
}
