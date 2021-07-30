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

type VM struct {
	DataPointer uint16
	Memory      [TAPE_LENGTH]byte
}

var VMBase VM = VM{DataPointer: 0, Memory: [TAPE_LENGTH]byte{}}

func (vm *VM) execute(out Writer, instruction Instruction) {
	switch instruction {
	case GO_LEFT:
		if vm.DataPointer == 0 {
			vm.DataPointer = TAPE_LENGTH - 1
		} else {
			vm.DataPointer -= 1
		}
	case GO_RIGHT:
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

func (vm *VM) RunCode(out Writer, instructions []Instruction, indexes []LoopIndexes) {
	length := len(instructions)

	for i := 0; i < length; i += 1 {
		switch instructions[i] {
		case LOOP_START:
			if vm.Memory[vm.DataPointer] == 0 {
				i = int(indexes[i])
			}
		case LOOP_END:
			if vm.Memory[vm.DataPointer] != 0 {
				i = int(indexes[i])
			}
		default:
			vm.execute(out, instructions[i])
		}
	}
}

func (vm *VM) RunString(str string, out Writer) {
	instructions, parseError, indexes, _ := Parse(str)
	StrictCheck(parseError)

	vm.RunCode(out, instructions, indexes)
}

func (vm *VM) RunFile(fileName string, out Writer) {
	file, fileError := ioutil.ReadFile(fileName)
	StrictCheck(fileError)

	vm.RunString(string(file), out)
}
