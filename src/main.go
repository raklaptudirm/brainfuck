package main

import (
	"fmt"
)

const TAPE_LENGTH uint16 = 30000

type Instruction uint8

const (
	GO_LEFT = iota
	GO_RIGHT
	INCREMENT
	DECREMENT
	INPUT
	OUTPUT
)

type VM struct {
	instructions []Instruction
	data_pointer uint64
	memory [TAPE_LENGTH]byte
}

func (vm VM) execute() {
	length := len(vm.instructions)

	for i := 0; i < length; i += 1 {
		switch vm.instructions[i] {
		case GO_LEFT:
			vm.data_pointer -= 1
		case GO_RIGHT:
			vm.data_pointer += 1
		case INPUT:
			fmt.Print("input: ")
			_, _ = fmt.Scanf("%d", &vm.memory[vm.data_pointer])
		case OUTPUT:
			fmt.Println(vm.memory[vm.data_pointer])
		case INCREMENT:
			vm.memory[vm.data_pointer] += 1
		case DECREMENT:
			vm.memory[vm.data_pointer] -= 1
		}
	}
}

func createVM (code string) *VM {
	bytecode := []Instruction{}
	length := len(code)

	for i := 0; i < length; i += 1 {
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
		}
	}

	vm := VM{instructions: bytecode, data_pointer: 0, memory: [30000]byte{}}
	return &vm
}

func main () {
	(*createVM(".>.,.+.-.")).execute()
}