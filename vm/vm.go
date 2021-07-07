package vm

import (
	"fmt"

	. "github.com/raklaptudirm/brainfuck/types"
)

type VM struct {
	DataPointer uint16
	Memory      [TAPE_LENGTH]byte
}

func (vm *VM) execute(instruction Instruction) {
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
		fmt.Print("input: ")
		_, _ = fmt.Scanf("%d", &vm.Memory[vm.DataPointer])
	case OUTPUT:
		fmt.Print(string(vm.Memory[vm.DataPointer]))
	case INCREMENT:
		vm.Memory[vm.DataPointer] += 1
	case DECREMENT:
		vm.Memory[vm.DataPointer] -= 1
	}
}

func (vm *VM) Run(instructions []Instruction) {
	length := len(instructions)

	for i := 0; i < length; i += 1 {
		switch instructions[i] {
		case LOOP_START:
			if vm.Memory[vm.DataPointer] == 0 {
				i += 1
				for loops := 0; instructions[i] != LOOP_END || loops != 0; i += 1 {
					switch instructions[i] {
					case LOOP_START:
						loops += 1
					case LOOP_END:
						loops -= 1
					}
				}
			}
		case LOOP_END:
			if vm.Memory[vm.DataPointer] != 0 {
				i -= 1
				for loops := 0; instructions[i] != LOOP_START || loops != 0; i -= 1 {
					switch instructions[i] {
					case LOOP_START:
						loops += 1
					case LOOP_END:
						loops -= 1
					}
				}
			}
		default:
			vm.execute(instructions[i])
		}
	}
}
