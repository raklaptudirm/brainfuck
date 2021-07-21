package vm

import (
	"fmt"
	. "io"

	. "github.com/raklaptudirm/brainfuck/types"
)

type VM struct {
	DataPointer uint16
	Memory      [TAPE_LENGTH]byte
}

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

func (vm *VM) Run(out Writer, instructions []Instruction, indexes []LoopIndexes) {
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
