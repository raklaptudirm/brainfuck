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
// Implementations can compile this intermediate code to any compilation
// target.
package instruction

import "fmt"

// Instruction represents a brainfuck instruction.
type Instruction interface {
	Instruction() string
}

// Value instruction changes the value of the current cell by X.
type Value struct {
	X byte
}

func (v *Value) Instruction() string {
	return fmt.Sprintf("%-16s %3d", "CHANGE_VAL", v.X)
}

// Pointer instruction changes the value pointer by X.
type Pointer struct {
	X int
}

func (p *Pointer) Instruction() string {
	return fmt.Sprintf("%-16s %d", "CHANGE_PTR", p.X)
}

// Input instruction takes a single byte as input from the user.
type Input struct{}

func (i *Input) Instruction() string {
	return "INPUT"
}

// Output instruction outputs the value of the current cell.
type Output struct{}

func (o *Output) Instruction() string {
	return "OUTPUT"
}

// StartLoop instruction signals the start of a loop.
type StartLoop struct{}

func (s *StartLoop) Instruction() string {
	return "START_LOOP"
}

// EndLoop instruction signals the end of a loop.
type EndLoop struct{}

func (e *EndLoop) Instruction() string {
	return "END_LOOP"
}

// Clear sets the current cell's value to 0.
type Clear struct{}

func (c *Clear) Instruction() string {
	return "CLEAR"
}
