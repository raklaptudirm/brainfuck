package parser

import (
	"errors"
	"fmt"

	. "github.com/raklaptudirm/brainfuck/types"
)

func Parse(code string) ([]Instruction, error, []LoopIndexes) {
	var parseError error = nil

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
			loops = append(loops, LoopIndexes(elements()))
		case "]":
			if len(loops) == 0 {
				parseError = errors.New(fmt.Sprintf("error %v:%v : Illeagal \"]\".", lines, column))
			} else {
				bytecode = append(bytecode, LOOP_END)

				loopStart := loops[last(loops)]

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
		parseError = errors.New("error: Unclosed \"[\".")
	}

	return bytecode, parseError, indexes
}
