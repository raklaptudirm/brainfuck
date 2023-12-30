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
	offset    int
}

// Finalize signals that the chunk has been built and no more instructions
// will be added. It will return a Chunk containing the built instructions.
// Calling Finalize when all loops have not been closed panics.
func (c *ChunkBuilder) Finalize() *Chunk {
	if !c.CanFinalize() {
		panic("chunk builder: can't finalize chunk because of unclosed loops")
	}

	// mark chunk as finalized
	c.finalized = true
	return &Chunk{ins: c.ins}
}

// IsFinalized informs whether the chunk has been finalized or not.
func (c *ChunkBuilder) IsFinalized() bool {
	return c.finalized
}

// CanFinalize informs whether Finalize can be called without panicking.
func (c *ChunkBuilder) CanFinalize() bool {
	return len(c.loopStack) == 0
}

// ChangeValue is a helper function for adding a Value instruction to the
// chunk, with the current offsets in mind.
func (c *ChunkBuilder) ChangeValue(by int8) {
	c.assertNotFinalized() // make sure chunk is not finalized
	c.optimizedPush(Value{X: byte(by), Offset: c.offset})
}

// ChangePointer is a helper function which represents adding a pointer
// instruction to the chunk but actually changes the offset.
func (c *ChunkBuilder) ChangePointer(change int) {
	c.assertNotFinalized() // make sure chunk is not finalized
	c.offset += change
}

// InputByte is a helper function for adding a Input instruction to the
// chunk, with the current offsets in mind.
func (c *ChunkBuilder) InputByte() {
	c.assertNotFinalized() // make sure chunk is not finalized
	c.optimizedPush(Input{Offset: c.offset})
}

// OutputByte is a helper function for adding a Output instruction to the
// chunk, with the current offsets in mind.
func (c *ChunkBuilder) OutputByte() {
	c.assertNotFinalized() // make sure chunk is not finalized
	c.push(Output{Offset: c.offset})
}

// StartLoop is a helper function for adding a StartLoop instruction to
// the chunk, with the current offsets in mind.
func (c *ChunkBuilder) StartLoop() {
	c.assertNotFinalized() // make sure chunk is not finalized

	c.loopStack = append(c.loopStack, len(c.ins)) // add to loop stack
	c.push(StartLoop{Offset: c.offset})           // push start loop
	c.offset = 0                                  // reset offset count
}

// EndLoop is a helper function which encapsulates adding a EndLoop
// instruction to the chunk.
func (c *ChunkBuilder) EndLoop() {
	c.assertNotFinalized() // make sure chunk is not finalized

	if len(c.loopStack) == 0 {
		panic("chunk builder: unexpected EndLoop")
	}

	last := len(c.loopStack) - 1     // last index of loopStack
	start := c.loopStack[last]       // last element of loopStack
	c.loopStack = c.loopStack[:last] // remove last element

	body := c.ins[start+1:]
	offset := c.ins[start].MemOffset()

	// remove loops which are never executed
	if c.isRedundantLoop(start, offset) {
		c.ins = c.ins[:start] // clear instruction slice
		c.offset = offset     // reset current offset
		return
	}

	// check if the loop body can be optimized
	if i, ok := optimizeLoopBody(body, offset, c.offset); ok {
		c.ins = c.ins[:start] // remove loop body
		c.put(i...)           // put optimized code

		// since the loop has been optimized, integrate it into the offset
		c.offset = offset
		return
	}

	// optimization failed, standard loop
	c.push(EndLoop{Offset: c.offset})
	c.offset = 0
}

// isRedundantLoop checks if a loop starting at the given position in the
// instruction chunk is redundant or not.
//
// Loops which are before any other instruction are redundant as all cells
// are 0 by default. Loops which start right after the end of another loop
// or a Clear instruction are redundant as the previous loop only exits
// when the cell is zero.
func (c *ChunkBuilder) isRedundantLoop(pos, offset int) bool {
	if pos == 0 {
		return true
	}
	
	ins := c.ins[pos-1]

	if ins.MemOffset() != offset {
		return false
	}

	switch v := ins.(type) {
	case Set:
		return v.X == 0
	case EndLoop:
		return true
	default:
		return false
	}
}

// assertNotFinalized makes sure that the chunk has not been finalized, and
// panics if it has been.
func (c *ChunkBuilder) assertNotFinalized() {
	if c.finalized {
		panic("chunk builder: chunk has already been finalized")
	}
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

// put puts the provided instructions into the chunk after optimizing them.
// Any calls to put after the chunk has been finalized panic.
func (c *ChunkBuilder) put(is ...Instruction) {
	for _, i := range is {
		c.optimizedPush(i)
	}
}

// optimizedPush adds the given instruction to the chunk after optimizing it.
// This function should not be exposed to external processes as some function
// calls may lead to unexpected results.
func (c *ChunkBuilder) optimizedPush(i Instruction) {

	// optimizations can only happen if the offsets are the same
	if c.last() != nil && c.last().MemOffset() == i.MemOffset() {
		switch curr := i.(type) {
		case Value:
			if curr.X == 0 {
				// Value instructions with X = 0 are redundant
				return
			}
			switch prev := c.last().(type) {
			// merge multiple Value instructions into a single one
			case Value:
				c.pop()
				if t := prev.X + curr.X; t != 0 {
					c.push(Value{X: t, Offset: curr.MemOffset()})
				}

				return

			// merge Value instructions into the Set instruction
			case Set:
				c.pop()
				c.push(Set{X: prev.X + curr.X, Offset: curr.MemOffset()})
				return
			}

		case Set, Input:
			// Clear and Input instructions make any adjacent Value, Clear, or Input
			// instructions redundant
			switch c.last().(type) {
			case Value, Set, Input:
				c.pop()
			}
		}
	}

	// push instruction into chunk
	c.push(i)
}

// push adds the given instruction to the chunk as given.
func (c *ChunkBuilder) push(i ...Instruction) {
	c.ins = append(c.ins, i...)
}

// optimizeLoopBody tries to optimize the given instructions which were
// found inside a loop. If successful, it returns the optimized
// instructions and true, other wise it returns nil and false.
func optimizeLoopBody(i []Instruction, start, end int) ([]Instruction, bool) {
	// any loop that can be optimized has to have a end offset of 0
	if end == 0 {
		switch len(i) {
		case 0:
			// empty loop
		case 1:
			// repeated changes to the value will just
			// loop until the current cell becomes 0
			if v, ok := i[0].(Value); ok {
				return []Instruction{Set{X: 0, Offset: start + v.Offset}}, true
			}
		default:
			// TODO: more loop optimizations
		}
	}

	// no optimizations found
	return nil, false
}
