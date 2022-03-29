// Copyright © 2021 Rak Laptudirm <raklaptudirm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"laptudirm.com/x/brainfuck/pkg/lexer"
	"laptudirm.com/x/brainfuck/pkg/parser"
	"laptudirm.com/x/brainfuck/pkg/token"
	"laptudirm.com/x/brainfuck/pkg/vm"
)

const Memory = 30000

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: brainfuck [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var l lexer.Lexer
	l.Init(string(file), handleErr)

	var p parser.Parser
	p.Init(&l, handleErr)

	program := p.ParseProgram()

	if l.ErrorCount > 0 || p.ErrorCount > 0 {
		os.Exit(1)
	}

	m := vm.New(Memory)
	err = m.Execute(program)
	if err != nil {
		printErr(err.Error())
		os.Exit(1)
	}
}

func handleErr(p token.Position, e string) {
	printErr("%s: %v\n", p, e)
}

func printErr(format string, a ...interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	return err
}
