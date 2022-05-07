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

// Package opcode implements the opcode compilation target. It provides
// structures and methods to compile an instruction.Chunk into opcode and
// to run the same.
package opcode

// Opcode represents a single opcode instruction.
type Opcode int

// Various opcode instructions.
const (
	_ Opcode = iota

	ChangeValue   // [code] [offset] [amount]
	ChangePointer // [code] [amount]
	InputByte     // [code] [offset]
	OutputByte    // [code] [offset]
	JumpIfZero    // [code] [offset] [jump-offset]
	JumpIfNotZero // [code] [jump-offset]
	SetValue      // [code] [offset] [amount]
)
