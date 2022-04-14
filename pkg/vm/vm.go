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

// Package vm implements a virtual machine that can run brainfuck bytecode.
package vm

import (
	"fmt"

	"laptudirm.com/x/brainfuck/pkg/bytecode"
)

// New creates a new virtualMachine with the provided tape length.
func New(size int) *virtualMachine {
	return &virtualMachine{tape: make([]byte, size), tapeSize: size}
}

// virtualMachine is a virtual machine.Brainfuck bytecode can be run on it
// by calling (virtualMachine).RunChunk.
type virtualMachine struct {
	tape     []byte // the vm's memory
	tapeSize int    // memory length
	pointer  int    // memory pointer
}

// InvalidBytecode is a wrapper error when the vm finds an invalid bytecode
// in the chunk.
type InvalidBytecode struct {
	Instruction bytecode.Instruction
}

// Error implements the error interface.
func (e *InvalidBytecode) Error() string {
	return fmt.Sprintf("vm: invalid bytecode %3d in chunk", e.Instruction)
}

// RunChunk runs a bytecode chunk in the virtual machine.
func (v *virtualMachine) RunChunk(c *bytecode.Chunk) error {
	for i, length := 0, c.Length(); i < length; i++ {
		// run the instruction
		switch c.Instruction(i) {
		case bytecode.IncreaseValue:
			v.tape[v.pointer]++

		case bytecode.DecreaseValue:
			v.tape[v.pointer]--

		case bytecode.IncreasePointer:
			v.pointer++

			// pointer rollover
			if v.pointer >= v.tapeSize {
				v.pointer = 0
			}

		case bytecode.DecreasePointer:
			v.pointer--

			// pointer rollover
			if v.pointer < 0 {
				v.pointer = v.tapeSize - 1
			}

		case bytecode.InputByte:
			fmt.Scanf("%c", &v.tape[v.pointer])

		case bytecode.OutputByte:
			fmt.Print(string(v.tape[v.pointer]))

		case bytecode.JumpIfZero:
			// check if value is zero
			if v.tape[v.pointer] == 0 {
				// jump
				i += int(c.Uint16(i + 1))
			}

			// jump over offset bytes
			i += 2

		case bytecode.JumpIfNotZero:
			// check if value is non-zero
			if v.tape[v.pointer] != 0 {
				// jump
				i -= int(c.Uint16(i + 1))
			}

			// jump over offset bytes
			i += 2

		default:
			// invalid instruction
			return &InvalidBytecode{c.Instruction(i)}
		}
	}

	return nil
}
