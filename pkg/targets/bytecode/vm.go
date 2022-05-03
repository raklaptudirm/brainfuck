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
	"io"
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
	// TODO: customizable memory length
	v := vm{memory: make([]byte, 30000)}
	length := len(c)

	// TODO: make this customizable
	buffer := printBuffer{
		writer: os.Stdout,
		aFlush: true,
		length: 50,
	}

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
			// syscalls are expensive, buffer output bytes
			buffer.Write(v.memory[v.pointer])

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

	// flush any remaining bytes
	buffer.Flush()
}

// printBuffer is a helper struct which buffers byte outputs for better
// performance, as syscalls are expensive.
type printBuffer struct {
	buffer []byte // backlog

	// options
	writer io.Writer // writer to output to
	aFlush bool      // automatically flush at intervals
	length int       // max backlog, only applicable if aFlush = true
}

// Write puts the given bytes into the backlog, and flushes it if it's
// length exceeds the provided maximum, and aFlush = true.
func (b *printBuffer) Write(bytes ...byte) {
	b.buffer = append(b.buffer, bytes...)

	if b.aFlush && len(b.buffer) > b.length {
		b.Flush()
	}
}

// Flush empties the backlog into the writer.
func (b *printBuffer) Flush() {
	// check if backlog is empty
	if len(b.buffer) == 0 {
		return
	}

	b.writer.Write(b.buffer)
	b.buffer = []byte{}
}
