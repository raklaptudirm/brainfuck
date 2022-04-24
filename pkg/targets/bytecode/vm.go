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

package bytecode

import (
	"fmt"
	"os"
)

// vm is a Virtual Machine which records the state of the brainfuck program
// as bytecode gets interpreted.
type vm struct {
	memory  []byte // memory tape
	pointer int    // memory pointer
}

// Run interprets a bytecode instruction slice. It panics if it encounters
// an unknown instruction. Currently the memory length of the vm is hard-
// coded to 30000.
func Run(c []byte) {
	v := vm{memory: make([]byte, 30000)}
	length := len(c)

	for i := 0; i < length; i++ {
		switch Bytecode(c[i]) {
		case ChangeValue:
			i++
			v.memory[v.pointer] += c[i]

		case ChangePointer:
			x, l := uintFromBytes(c[i+1:])
			v.pointer += int(x)
			i += l

		case InputByte:
			fmt.Scanf("%c", &v.memory[v.pointer])
		case OutputByte:
			os.Stdout.Write([]byte{v.memory[v.pointer]})

		case JumpIfZero:
			x, l := uintFromBytes(c[i+1:])
			i += l // jump over offset bytes

			if v.memory[v.pointer] == 0 {
				// current cell 0, so jump to loop end
				i += int(x)
			}
		case JumpIfNotZero:
			x, l := uintFromBytes(c[i+1:])
			if v.memory[v.pointer] == 0 {
				// current cell 0, so jump over offset bytes
				i += l
			} else {
				// current cell not 0, so jump back to loop start
				i -= int(x)
			}

		case ClearValue:
			v.memory[v.pointer] = 0

		default:
			// invalid bytecode instruction
			panic(fmt.Sprintf("vm: invalid bytecode instruction %2x", c[i]))
		}
	}
}
