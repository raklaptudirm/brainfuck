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

// Package bytecode defines an bytecode instruction set used to represent
// brainfuck code and related functions.
package bytecode

import "fmt"

// Chunk represents a chunk of bytecode.
type Chunk struct {
	Name string // name of the chunk
	code []Instruction
}

// Write writes a new instruction to the bytecode chunk.
func (c *Chunk) Write(i Instruction) {
	c.code = append(c.code, i)
}

// Length returns the length of the current chunk.
func (c *Chunk) Length() int {
	return len(c.code)
}

// Instruction returns the instruction as offset i in the chunk.
func (c *Chunk) Instruction(i int) Instruction {
	return c.code[i]
}

// String disassembles the chunk into a human readable string.
func (c *Chunk) String() string {
	// chunk header
	s := fmt.Sprintf("== %s [%d bytes] ==\n", c.Name, len(c.code))

	// disassemble each instruction
	for i, next := 0, 0; i < len(c.code); i = next {
		var ins string
		ins, next = c.disassembleInstruction(i)
		s += ins
	}

	return s
}

// disassembleInstruction converts the instruction at the given offset in
// the current chunk into a human readable string with a trailing newline.
func (c *Chunk) disassembleInstruction(offset int) (string, int) {
	switch i := c.code[offset]; i {
	// simple instructions
	case IncreaseValue, DecreaseValue, IncreasePointer, DecreasePointer, InputByte, OutputByte:
		return fmt.Sprintf("%4d %s\n", offset, i), offset + 1

	// instructions with one argument
	case JumpIfZero, JumpIfNotZero:
		return fmt.Sprintf("%4d %-17s %3d\n", offset, i, byte(c.code[offset+1])), offset + 2

	// invalid instruction
	default:
		return fmt.Sprintf("%4d %d\n", offset, i), offset + 1
	}
}

// Instruction represents a single brainfuck bytecode instruction.
type Instruction byte

// set of constants representing various brainfuck bytecode instructions.
const (
	IncreaseValue Instruction = iota
	DecreaseValue
	IncreasePointer
	DecreasePointer
	InputByte
	OutputByte
	JumpIfZero
	JumpIfNotZero
)

var instructions = [...]string{
	IncreaseValue:   "INCREASE_VAL",
	DecreaseValue:   "DECREASE_VAL",
	IncreasePointer: "INCREASE_PTR",
	DecreasePointer: "DECREASE_PTR",
	InputByte:       "INPUT_BYTE",
	OutputByte:      "OUTPUT_BYTE",
	JumpIfZero:      "JUMP_IF_ZERO",
	JumpIfNotZero:   "JUMP_IF_NOT_ZERO",
}

// String converts a brainfuck bytecode instruction into a human-readable
// string.
func (i Instruction) String() string {
	if int(i) < len(instructions) {
		return instructions[i]
	}

	return "ILLEGAL"
}
