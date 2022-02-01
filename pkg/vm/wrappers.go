package vm

import (
	"io"
	"io/ioutil"

	"github.com/raklaptudirm/brainfuck/pkg/errors"
	"github.com/raklaptudirm/brainfuck/pkg/parser"
)

// RunString runs a given brainfuck source string,
// by parsing it and then using VM.RunCode.
func (vm *VM) RunString(str string, out io.Writer) {
	bytecode, parseError, _ := parser.Parse(str)
	errors.StrictCheck(parseError)

	vm.RunCode(out, bytecode)
}

// RunFile runs a brainfuck source file, by reading it,
// parsing it, and then using VM.RunCode.
func (vm *VM) RunFile(fileName string, out io.Writer) {
	file, fileError := ioutil.ReadFile(fileName)
	errors.StrictCheck(fileError)

	vm.RunString(string(file), out)
}
