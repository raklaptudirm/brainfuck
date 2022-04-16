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
// token stream into an abstract syntax tree.
package parser

import (
	"fmt"

	"laptudirm.com/x/brainfuck/pkg/ast"
	"laptudirm.com/x/brainfuck/pkg/token"
)

// Parse parses a brainfuck token stream into an abstract syntax tree.
func Parse(tokens <-chan token.Token) (*ast.Program, error) {
	p := parser{tokens: tokens}
	return p.program()
}

// parser is a state machine which represents the current parsing state.
type parser struct {
	tokens  <-chan token.Token // token stream
	current token.Token        // current token
	tokType token.Type         // type of current
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
func (p *parser) program() (*ast.Program, error) {
	program := &ast.Program{}

	for p.next(); p.tokType != token.Eof; p.next() {
		operation, err := p.operation()
		if err != nil {
			return nil, err
		}

		program.Operations = append(program.Operations, operation)
	}

	return program, nil
}

// operation parses a brainfuck operation from the token stream.
func (p *parser) operation() (ast.Operation, error) {
	switch p.tokType {
	// simple operations
	case token.Plus, token.Minus, token.LeftArrow, token.RightArrow, token.Comma, token.Period:
		operator := ast.Operator(p.current)
		return &operator, nil

	// loop operation
	case token.LeftBracket:
		return p.loop()

	// invalid
	case token.RightBracket:
		return nil, &SyntaxError{Token: p.current, message: ErrNotOpened}
	default:
		panic("parser: invalid token in token stream")
	}
}

// loop parses a brainfuck loop from the token stream.
func (p *parser) loop() (*ast.Loop, error) {
	loop := &ast.Loop{}

	for p.next(); p.tokType != token.RightBracket; p.next() {
		if p.tokType == token.Eof {
			return nil, &SyntaxError{Token: p.current, message: ErrNotClosed}
		}

		operation, err := p.operation()
		if err != nil {
			return nil, err
		}

		loop.Operations = append(loop.Operations, operation)
	}

	return loop, nil
}

// next gets the next token from the token stream and stores it in current.
func (p *parser) next() {
	p.current = <-p.tokens
	p.tokType = p.current.Type
}
