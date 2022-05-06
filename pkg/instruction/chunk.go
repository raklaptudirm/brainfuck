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

package instruction

import "fmt"

// Chunk represents an immutable list of instructions.
type Chunk struct {
	ins []Instruction
}

// String converts a Chunk into a human readable string.
func (c *Chunk) String() string {
	var s string
	length := c.Len()

	for i := 0; i < length; i++ {
		s += fmt.Sprintf("%4d %s\n", i, c.Instruction(i).Instruction())
	}

	return s
}

// Instruction fetches the ith instruction in the Chunk. It will panic if i
// is greater than the length of the Chunk.
func (c *Chunk) Instruction(i int) Instruction {
	return c.ins[i]
}

// Len returns the length of the Chunk.
func (c *Chunk) Len() int {
	return len(c.ins)
}
