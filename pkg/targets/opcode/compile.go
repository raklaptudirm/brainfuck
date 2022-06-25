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
	"reflect"

	"laptudirm.com/x/brainfuck/pkg/instruction"
)

// Compile compiles an instruction.Chunk into opcode, which is represented by
// a slice of integers.
func Compile(c *instruction.Chunk) []int {
	var dst []int   // result slice
	var stack []int // loop stack

	length := c.Len()
	for i := 0; i < length; i++ {
		ins := c.Instruction(i)

		switch v := ins.(type) {
		case instruction.Value:
			// [code] [offset] [amount]
			dst = append(dst, int(ChangeValue), v.Offset, int(v.X))

		case instruction.Input:
			// [code] [offset]
			dst = append(dst, int(InputByte), v.Offset)

		case instruction.Output:
			// [code] [offset]
			dst = append(dst, int(OutputByte), v.Offset)

		case instruction.StartLoop:
			// [code] [offset] [jump-offset]
			dst = append(dst, int(JumpIfZero), v.Offset, 0)
			stack = append(stack, len(dst))

		case instruction.EndLoop:
			if len(stack) == 0 {
				// no loops opened, unreachable
				panic("opcode: compile: unexpected EndLoop instruction in chunk")
			}

			// [code] [offset] [jump-offset]
			dst = append(dst, int(JumpIfNotZero), v.Offset, 0)

			start := stack[len(stack)-1] // get loop start index
			stack = stack[:len(stack)-1] // pop loop index

			// difference between loop start and end
			diff := len(dst) - start

			// backpatch jump-offsets
			dst[len(dst)-1], dst[start-1] = diff, diff

		case instruction.Set:
			// [code] [offset] [amount]
			dst = append(dst, int(SetValue), v.Offset, int(v.X))

		default:
			// unreachable
			t := reflect.ValueOf(ins).Elem().Type() // get instruction type
			panic(fmt.Sprintf("opcode: compile: invalid instruction type %s in chunk", t))
		}
	}

	if len(stack) > 0 {
		// unclosed loops, unreachable
		panic("opcode: compile: unexpected end of chunk, unpaired StartLoop instructions")
	}

	return dst
}
