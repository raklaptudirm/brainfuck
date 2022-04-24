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

// Package bytecode implements the bytecode compilation target. It provides
// structures and methods to compile an instruction.Chunk into bytecode and
// to run the same.
package bytecode

// Bytecode represents a single bytecode instruction.
type Bytecode byte

// constants representing various bytecode instructions.
const (
	ChangeValue Bytecode = iota
	ChangePointer
	InputByte
	OutputByte
	JumpIfZero
	JumpIfNotZero
	ClearValue
)
