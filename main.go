// Copyright © 2022 Rak Laptudirm <raklaptudirm@gmail.com>
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
	"laptudirm.com/x/brainfuck/pkg/targets/opcode"
)

func main() {
	if err := mainFunc(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mainFunc() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: brainfuck <file>")
	}

	// extract filename
	filename := os.Args[1]

	// read source code
	source, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// parse source code
	ins, err := parser.Parse(lexer.Lex(source))
	if err != nil {
		return err
	}

	// compile to opcode and run
	oc := opcode.Compile(ins)
	opcode.Run(oc)

	return nil
}
