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

// Package compiler implements a brainfuck compiler which compiles an
// abstract syntax tree into a bytecode chunk.
package compiler

import (
	"fmt"
	"reflect"

	"laptudirm.com/x/brainfuck/pkg/ast"
	"laptudirm.com/x/brainfuck/pkg/bytecode"
	"laptudirm.com/x/brainfuck/pkg/token"
)

// Compile compiles the given abstract syntax tree into a bytecode chunk.
func Compile(p *ast.Program) (*bytecode.Chunk, error) {
	c := compiler{&bytecode.Chunk{}}
	if err := c.program(p); err != nil {
		return nil, err
	}

	return c.chunk, nil
}

// compiler is a state machine which keeps track of the chunk into which
// the abstract syntax tree is getting compiled.
type compiler struct {
	chunk *bytecode.Chunk
}

// program compiles an ast.Program into the chunk.
func (c *compiler) program(p *ast.Program) error {
	// compile program operations
	for _, op := range p.Operations {
		if err := c.operation(op); err != nil {
			return err
		}
	}

	return nil
}

// operation compiles an ast.Operation node into the chunk.
func (c *compiler) operation(node ast.Operation) error {
	switch op := node.(type) {
	case *ast.Loop:
		return c.loop(op)
	case *ast.Operator:
		return c.operator(op)
	default:
		return fmt.Errorf("compiler: unexpected ast node of type %s", reflect.TypeOf(op))
	}
}

// loop compiles an ast.Loop node into the chunk.
func (c *compiler) loop(l *ast.Loop) error {
	c.chunk.Write(bytecode.JumpIfZero) // loop start instruction
	start := c.chunk.Length()          // start offset
	c.chunk.WriteUint16(0xFFFF)        // write placeholder

	// compile loop operations
	for _, op := range l.Operations {
		if err := c.operation(op); err != nil {
			return err
		}
	}

	c.chunk.Write(bytecode.JumpIfNotZero) // loop end instruction
	end := c.chunk.Length()               // end offset
	c.chunk.WriteUint16(0xFFFF)           // write placeholder

	diff := end - start // difference between offsets

	// backpatch offsets
	c.chunk.WriteUint16At(start, uint16(diff+1))
	c.chunk.WriteUint16At(end, uint16(diff-1))

	return nil
}

// operator compiles an ast.Operator node into the chunk.
func (c *compiler) operator(p *ast.Operator) error {
	switch t := token.Token(*p).Type; t {
	case token.Plus:
		c.chunk.Write(bytecode.IncreaseValue)
	case token.Minus:
		c.chunk.Write(bytecode.DecreaseValue)
	case token.LeftArrow:
		c.chunk.Write(bytecode.IncreasePointer)
	case token.RightArrow:
		c.chunk.Write(bytecode.DecreasePointer)
	case token.Comma:
		c.chunk.Write(bytecode.InputByte)
	case token.Period:
		c.chunk.Write(bytecode.OutputByte)
	default:
		return fmt.Errorf("compiler: unexpected ast operator node with %s token", t)
	}

	return nil
}
