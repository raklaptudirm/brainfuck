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

package opcode

import (
	"fmt"
	"io"
	"os"
)

// vm is a Virtual Machine which records the state of the brainfuck program
// as opcode gets interpreted.
type vm struct {
	memory  []byte // memory tape
	pointer int    // memory pointer
}

func Run(oc []int) {
	// TODO: make these options customizable
	v := vm{memory: make([]byte, 30000)}
	buffer := printBuffer{
		writer:    os.Stdout,
		autoFlush: true,
		length:    50,
	}

	length := len(oc)
	for i := 0; i < length; i++ {
		switch Opcode(oc[i]) {
		case ChangeValue:
			pointer := v.pointer + oc[i+1]     // calculate pointer offset
			v.memory[pointer] += byte(oc[i+2]) // change value by amount
			i += 2                             // update instruction pointer

		case ChangePointer:
			i++                // update instruction pointer
			v.pointer += oc[i] // update memory pointer

		case InputByte:
			i++                                 // update instruction pointer
			pointer := v.pointer + oc[i]        // calculate pointer offset
			fmt.Scanf("%c", &v.memory[pointer]) // store input in memory

		case OutputByte:
			i++                             // update instruction pointer
			pointer := v.pointer + oc[i]    // calculate pointer offset
			buffer.Write(v.memory[pointer]) // output current cell value

		case JumpIfZero:
			i++                // update instruction pointer
			v.pointer += oc[i] // change pointer by offset

			i++           // update instruction pointer
			jump := oc[i] // get jump offset

			// jump if zero
			if v.memory[v.pointer] == 0 {
				i += jump
			}

		case JumpIfNotZero:
			i++           // update instruction pointer
			jump := oc[i] // get jump offset

			// jump back if not zero
			if v.memory[v.pointer] != 0 {
				i -= jump
			}

		case SetValue:
			i++                          // update instruction pointer
			pointer := v.pointer + oc[i] // calculate pointer offset

			i++                  // update instruction pointer
			value := byte(oc[i]) // get set value

			v.memory[pointer] = value // clear current cell

		default:
			panic(fmt.Sprintf("opcode: run: invalid opcode %x", uint(oc[i])))
		}
	}

	// flush any remaining output
	buffer.Flush()
}

// printBuffer is a helper struct which buffers byte outputs for better
// performance, as syscalls are expensive.
type printBuffer struct {
	buffer []byte // backlog

	// options
	writer    io.Writer // writer to output to
	autoFlush bool      // automatically flush at intervals
	length    int       // max backlog, only applicable if aFlush = true
}

// Write puts the given bytes into the backlog, and flushes it if it's
// length exceeds the provided maximum, and aFlush = true.
func (b *printBuffer) Write(bytes ...byte) {
	b.buffer = append(b.buffer, bytes...)

	if b.autoFlush && len(b.buffer) > b.length {
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
