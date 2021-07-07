package parser

import (
	"errors"
	"fmt"

	. "github.com/raklaptudirm/brainfuck/types"
)

func Parse(code string) ([]Instruction, error) {
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
