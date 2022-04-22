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

// Package parser implements a brainfuck parser that parses a brainfuck
// token stream into an instruction set.
package parser

import (
	"fmt"

	"laptudirm.com/x/brainfuck/pkg/instruction"
	"laptudirm.com/x/brainfuck/pkg/token"
)

// Parse parses a brainfuck token stream into an abstract syntax tree.
func Parse(tokens <-chan token.Token) (*instruction.Chunk, error) {
	p := parser{tokens: tokens}
	return p.program()
}

// parser is a state machine which represents the current parsing state.
type parser struct {
	tokens  <-chan token.Token // token stream
	current token.Token        // current token
}

// SyntaxError represents a brainfuck syntax error at a particular token.
type SyntaxError struct {
	Token   token.Token // token at which error occurred
	message error       // the error
}

// Error implements the error interface.
func (e *SyntaxError) Error() string {
	return fmt.Sprintf("parser: %s: %v", e.Token.Position, e.message)
}

// Unwrap exposes the underlying error in SyntaxError.
func (e *SyntaxError) Unwrap() error {
	return e.message
}

// Error values which are held inside SyntaxError.
var (
	ErrNotOpened = fmt.Errorf("unexpected token ']', no open loop")
	ErrNotClosed = fmt.Errorf("unexpected token EOF, loop not closed")
)

// program parses a brainfuck program from the token stream.
func (p *parser) program() (*instruction.Chunk, error) {
	var c instruction.ChunkBuilder
	var stack []token.Token // loop stack

parseLoop:
	for {
		switch p.next(); p.current.Type {

		// end of token stream
		case token.Eof:
			break parseLoop

		// value changing commands
		case token.Plus:
			c.Put(&instruction.Value{X: 1})
		case token.Minus:
			c.Put(&instruction.Value{X: 255}) // -1 mod 256(byte)

		// pointer changing commands
		case token.LeftArrow:
			c.Put(&instruction.Pointer{X: -1})
		case token.RightArrow:
			c.Put(&instruction.Pointer{X: 1})

		// i/o commands
		case token.Comma:
			c.Put(&instruction.Input{})
		case token.Period:
			c.Put(&instruction.Output{})

		// looping constructs
		case token.LeftBracket:
			stack = append(stack, p.current) // push
			c.Put(&instruction.StartLoop{})
		case token.RightBracket:
			if len(stack) == 0 {
				// no opened loop, syntax error
				return nil, &SyntaxError{p.current, ErrNotOpened}
			}

			stack = stack[:len(stack)-1] // pop
			c.Put(&instruction.EndLoop{})

		default:
			// unreachable
			panic("parser: invalid token from scanner")
		}
	}

	// check for unclosed loops
	if len(stack) > 0 {
		return nil, &SyntaxError{stack[len(stack)-1], ErrNotClosed}
	}

	// finalize and return chunk
	return c.Finalize(), nil
}

// next gets the next token from the token stream and stores it in current.
func (p *parser) next() {
	p.current = <-p.tokens
}
