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

// Package instruction defines an optimized brainfuck instruction set which
// acts as an intermediate representation for brainfuck code.
//
// Implementations can compile this intermediate code to any compilation
// target. All conforming processes must properly handle all instructions
// defined in this package, and reject any others. The interpretetion of
// an instruction must also follow it's definition.
//
// Since an instruction chunk can only be created by the builder,
// implementations may assume any *instruction.Chunk to be valid, i.e,
// composed only of the specified instructions and with properly matched
// loops.
package instruction

import "fmt"

// Instruction represents a brainfuck instruction.
type Instruction interface {
	Instruction() string
	MemOffset() int
}

// Value instruction changes the value of the cell at the given offset from
// the current cell by X.
type Value struct {
	X      byte
	Offset int
}

func (v *Value) Instruction() string {
	return fmt.Sprintf("Change Value at %d by %d", v.Offset, int8(v.X))
}

func (v *Value) MemOffset() int {
	return v.Offset
}

// Pointer instruction changes the pointer by X.
type Pointer struct {
	X int
}

func (p *Pointer) Instruction() string {
	return fmt.Sprintf("Change Pointer by %d", p.X)
}

func (p *Pointer) MemOffset() int {
	return p.X
}

// Input instruction takes a single byte as input from the user and stores
// it in the cell at the given offset from the current cell.
type Input struct {
	Offset int
}

func (i *Input) Instruction() string {
	return fmt.Sprintf("Input Byte at %d", i.Offset)
}

func (i *Input) MemOffset() int {
	return i.Offset
}

// Output instruction outputs the value of the cell at the given offset
// from the current cell as a string, i.e. 65 -> A.
type Output struct {
	Offset int
}

func (o *Output) Instruction() string {
	return fmt.Sprintf("Output Byte at %d", o.Offset)
}

func (o *Output) MemOffset() int {
	return o.Offset
}

// StartLoop instruction signals the start of a loop, after moving the
// pointer by the given offset.
type StartLoop struct {
	Offset int
}

func (s *StartLoop) Instruction() string {
	return fmt.Sprintf("Start Loop at %d", s.Offset)
}

func (s *StartLoop) MemOffset() int {
	return s.Offset
}

// EndLoop instruction signals the end of a loop.
type EndLoop struct{}

func (e *EndLoop) Instruction() string {
	return "End Loop"
}

func (e *EndLoop) MemOffset() int {
	return 0
}

// Set sets value of the cell at the given offset from the current cell to
// the given value.
type Set struct {
	X      byte
	Offset int
}

func (c *Set) Instruction() string {
	return fmt.Sprintf("Set %d at %d", c.X, c.Offset)
}

func (c *Set) MemOffset() int {
	return c.Offset
}
