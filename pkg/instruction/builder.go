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

// ChunkBuilder is helper struct which is used to build an optimized
// instruction Chunk. It's zero value is safe to use.
type ChunkBuilder struct {
	ins       []Instruction
	loopStack []int
	finalized bool
}

// Finalize signals that the chunk has been built and no more instructions
// will be added. It will return a Chunk containing the built instructions.
// Calling Finalize when all loops have not been closed panics.
func (c *ChunkBuilder) Finalize() *Chunk {
	if !c.CanFinalize() {
		panic("chunk: can't finalize chunk because of unclosed loops")
	}

	// mark chunk as finalized
	c.finalized = true
	return NewChunk(c.ins)
}

// Put puts the provided instructions into the chunk after optimizing them.
// Any calls to Put after the chunk has been finalized panic.
func (c *ChunkBuilder) Put(is ...Instruction) {
	if c.IsFinalized() {
		// can't push to finalized chunk
		panic("chunk: chunk has already been finalized")
	}

	for _, i := range is {
		c.push(i)
	}
}

// IsFinalized informs whether the chunk has been finalized or not.
func (c *ChunkBuilder) IsFinalized() bool {
	return c.finalized
}

// CanFinalize informs whether Finalize can be called without panicking.
func (c *ChunkBuilder) CanFinalize() bool {
	return len(c.loopStack) == 0
}

// last is syntactic sugar for getting the last item in the instructions.
func (c *ChunkBuilder) last() Instruction {
	if len(c.ins) == 0 {
		return nil
	}

	return c.ins[len(c.ins)-1]
}

// pop is syntactic sugar for removing the last instruction.
func (c *ChunkBuilder) pop() {
	if len(c.ins) == 0 {
		return
	}

	c.ins = c.ins[:len(c.ins)-1]
}

// push adds the given instruction to the chunk after optimizing it.
func (c *ChunkBuilder) push(i Instruction) {
	switch ins := i.(type) {
	case *Value:
		if ins.X == 0 {
			// Value instructions with X = 0 are redundant
			return
		}

		// if last instruction is also a Value, merge with it
		if v, ok := c.last().(*Value); ok {
			c.pop()

			if t := v.X + ins.X; t == 0 {
				// X = 0
				return
			} else {
				i = &Value{X: t}
			}
		}

	case *Pointer:
		if ins.X == 0 {
			// Pointer instructions with X = 0 are redundant
			return
		}

		// if last instruction is also a Pointer, merge with it
		if p, ok := c.last().(*Pointer); ok {
			c.pop()

			if t := p.X + ins.X; t == 0 {
				// X = 0
				return
			} else {
				i = &Pointer{X: t}
			}
		}

	case *Clear:
		// Clear instruction makes any previous Value instructions redundant
		if _, ok := c.last().(*Value); ok {
			c.pop()
		}

		// multiple Clear instructions are redundant
		if _, ok := c.last().(*Clear); ok {
			return
		}

	case *StartLoop:
		// add to loop stack
		c.loopStack = append(c.loopStack, len(c.ins))

	case *EndLoop:
		if len(c.loopStack) == 0 {
			panic("chunk: unexpected *EndLoop")
		}

		last := len(c.loopStack) - 1     // last index of loopStack
		start := c.loopStack[last]       // last element of loopStack
		c.loopStack = c.loopStack[:last] // remove last element

		// check if the loop body can be optimized
		if i, ok := optimizeLoopBody(c.ins[start+1:]); ok {
			c.ins = c.ins[:start] // remove loop body
			c.Put(i...)           // put optimized code
			return
		}
	}

	c.ins = append(c.ins, i)
}

// optimizeLoopBody tries to optimize the given instructions which were
// found inside a loop. If successful, it returns the optimized
// instructions and true, other wise it returns nil and false.
func optimizeLoopBody(i []Instruction) ([]Instruction, bool) {
	switch len(i) {
	case 0:
		// empty loop
	case 1:
		// repeated changes to the value will just
		// loop until the current cell becomes 0
		if _, ok := i[0].(*Value); ok {
			return []Instruction{&Clear{}}, true
		}
	default:
		// TODO: more loop optimizations
	}

	// no optimizations found
	return nil, false
}
