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
			} else {
				loops -= 1
				bytecode = append(bytecode, LOOP_END)
			}
		case "\n":
			lines += 1
			column = 0
		}
	}

	if loops != 0 {
		parseError = errors.New("error: Unclosed \"[\".")
	}

	return bytecode, parseError
}

func strictCheck(e error) {
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}
}

func check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}

func main() {
	brainfuck := VIRTUAL_MACHINE
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Welcome to brainfuck v1.0.0.")
		for {
			fmt.Print("> ")
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				break
			} else {
				instructions, parseError := parse(input)

				if check(parseError) {
					brainfuck.run(instructions)
				}

				length := len(instructions)
				for i := 0; i < length; i++ {
					if instructions[i] == OUTPUT {
						fmt.Print("\n")
					}
				}
			}
		}
	} else {
		file, fileError := ioutil.ReadFile(args[0])
		strictCheck(fileError)

		instructions, parseError := parse(string(file))
		strictCheck(parseError)

		brainfuck.run(instructions)
	}
}
