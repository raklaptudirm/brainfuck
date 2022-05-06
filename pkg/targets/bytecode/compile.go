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
	"laptudirm.com/x/brainfuck/pkg/instruction"
)

// Compile compiles an instruction.Chunk into bytecode. Compile assumes that
// the provided chunk is a set of valid and optimized instructions.
func Compile(c *instruction.Chunk) []byte {
	var dst []byte  // result slice
	var stack []int // loop stack

	length := c.Len()
	for i := 0; i < length; i++ {
		ins := c.Instruction(i)

		switch v := ins.(type) {
		case *instruction.Value:
			dst = append(dst, byte(ChangeValue), v.X)
			dst = append(dst, uintToBytes(uint64(v.Offset))...)

		case *instruction.Pointer:
			dst = append(dst, byte(ChangePointer))
			dst = append(dst, uintToBytes(uint64(v.X))...)

		case *instruction.Input:
			dst = append(dst, byte(InputByte))
			dst = append(dst, uintToBytes(uint64(v.Offset))...)
		case *instruction.Output:
			dst = append(dst, byte(OutputByte))
			dst = append(dst, uintToBytes(uint64(v.Offset))...)

		case *instruction.StartLoop:
			dst = append(dst, byte(JumpIfZero))
			dst = append(dst, uintToBytes(uint64(v.Offset))...)
			stack = append(stack, len(dst)) // push index to stack
		case *instruction.EndLoop:
			start := stack[len(stack)-1] // get loop start index
			stack = stack[:len(stack)-1] // pop loop index

			// difference between loop start and end
			diff := len(dst) - start + 1

			closeOffset := uintToBytes(uint64(diff))                   // JumpIfNotZero offset
			openOffset := uintToBytes(uint64(diff + len(closeOffset))) // JumpIfZero offset

			body := make([]byte, len(dst)-start)
			copy(body, dst[start:]) // copy loop body

			dst = append(dst[:start], openOffset...) // backpatch jump offset
			dst = append(dst, body...)               // re-append loop body

			dst = append(dst, byte(JumpIfNotZero))
			dst = append(dst, closeOffset...) // write offset

		case *instruction.Clear:
			dst = append(dst, byte(ClearValue))
			dst = append(dst, uintToBytes(uint64(v.Offset))...)
		}
	}

	return dst
}

// singleByteLim is the maximum number that can be represented by a single
// byte under the bytecode number encoding system.
const singleByteLim = 247

// uintToBytes converts a uint64 into a byte slice following the bytecode
// number encoding system.
//
// If the number is less than or equal to 247(singleByteLim), it is encoded
// as a single byte equal to it's value.
//
// Otherwise, the number is encoded as [magic] [number bytes], where:
// magic: a magic identifier byte which is equal to 247 + len(number bytes),
//   where 247 is the singleByteLim.
// number bytes: the number encoded in big endian format.
func uintToBytes(u uint64) []byte {
	var bytes int

	// TODO: check if this can be made better
	switch {
	case u <= singleByteLim:
		return []byte{byte(u)}
	case u < 1<<8:
		bytes = 1
	case u < 1<<16:
		bytes = 2
	case u < 1<<24:
		bytes = 3
	case u < 1<<32:
		bytes = 4
	case u < 1<<40:
		bytes = 5
	case u < 1<<48:
		bytes = 6
	case u < 1<<56:
		bytes = 7
	default:
		bytes = 8
	}

	const mask = 1<<8 - 1

	buf := make([]byte, bytes+1)
	buf[0] = byte(singleByteLim + bytes)

	for i := 0; i < bytes; i++ {
		buf[bytes-i] = byte(u & mask)
		u >>= 8
	}

	return buf
}

// uintFromBytes reads a number encoded in the bytecode number encoding
// system from the byte slice and returns it's value, along with the number
// of bytes it occupies. See the documentation of uintToBytes for the
// specification of the encoding system.
func uintFromBytes(b []byte) (uint64, int) {
	if b[0] <= singleByteLim {
		// single byte number
		return uint64(b[0]), 1
	}

	// get number of bytes from magic number
	bytes := int(b[0] - singleByteLim)

	var n uint64
	for i := 1; i <= bytes; i++ {
		n <<= 8
		n |= uint64(b[i])
	}

	return n, bytes + 1
}
