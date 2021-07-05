package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const TAPE_LENGTH uint16 = 30000

var VIRTUAL_MACHINE VM = VM{dataPointer: 0, memory: [TAPE_LENGTH]byte{}}

type Instruction uint8

const (
	GO_LEFT = iota
	GO_RIGHT
	INCREMENT
	DECREMENT
	INPUT
	OUTPUT
	LOOP_START
	LOOP_END
)

type VM struct {
	dataPointer uint64
	memory      [TAPE_LENGTH]byte
}

func (vm *VM) execute(instruction Instruction) {
	switch instruction {
	case GO_LEFT:
		vm.dataPointer -= 1
	case GO_RIGHT:
		vm.dataPointer += 1
	case INPUT:
		fmt.Print("input: ")
		_, _ = fmt.Scanf("%d", &vm.memory[vm.dataPointer])
	case OUTPUT:
		fmt.Print(string(vm.memory[vm.dataPointer]))
	case INCREMENT:
		vm.memory[vm.dataPointer] += 1
	case DECREMENT:
		vm.memory[vm.dataPointer] -= 1
	}
}

func (vm *VM) run(instructions []Instruction) {
	length := len(instructions)

	for i := 0; i < length; i += 1 {
		switch instructions[i] {
		case LOOP_START:
			if vm.memory[vm.dataPointer] == 0 {
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
			if vm.memory[vm.dataPointer] != 0 {
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

func parse(code string) ([]Instruction, error) {
	var parseError error = nil

	bytecode := []Instruction{}
	length := len(code)
	loops := 0
	lines := 1
	column := 0

	for i := 0; i < length; i += 1 {
		column += 1

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
			loops += 1
			bytecode = append(bytecode, LOOP_START)
		case "]":
			if loops == 0 {
				parseError = errors.New(fmt.Sprintf("error %v:%v : Illeagal \"]\".", lines, column))
			}

			loops -= 1
			bytecode = append(bytecode, LOOP_END)
		case "\n":
			lines += 1
			column = 0
		}
	}

	if loops != 0 {
		parseError = errors.New("error: Unclosed \"[\" (LOOP_START)")
	}

	return bytecode, parseError
}

func check(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
}

func main() {
	brainfuck := VIRTUAL_MACHINE
	args := os.Args[1:]

	file, fileError := ioutil.ReadFile(args[0])
	check(fileError)

	instructions, parseError := parse(string(file))
	check(parseError)

	brainfuck.run(instructions)
}
