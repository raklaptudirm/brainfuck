// Copyright © 2021 Rak Laptudirm <raklaptudirm@gmail.com>
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

package parser

import (
	"github.com/raklaptudirm/brainfuck/pkg/ast"
	"github.com/raklaptudirm/brainfuck/pkg/lexer"
	"github.com/raklaptudirm/brainfuck/pkg/token"
)

type Parser struct {
	l *lexer.Lexer

	err lexer.ErrorHandler

	tok token.Token
	pos token.Position
	lit string

	ErrorCount int
}

func (p *Parser) Init(l *lexer.Lexer, err lexer.ErrorHandler) {
	p.l = l
	p.err = err
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.next(); p.tok != token.EOF; p.next() {
		op := p.parseOperation()
		program.Operations = append(program.Operations, op)
	}

	return program
}

func (p *Parser) parseOperation() ast.Operation {
	switch p.tok {
	case token.COMMENT:
		return &ast.Comment{Literal: p.lit}
	case token.SLOOP:
		return p.parseLoop()
	case token.ELOOP:
		p.error(p.pos, "unexpected ]")
		return nil
	default:
		return &ast.Operator{Token: p.tok}
	}
}

func (p *Parser) parseLoop() *ast.Loop {
	loop := &ast.Loop{}
	pos := p.pos

	for p.next(); p.tok != token.ELOOP && p.tok != token.EOF; p.next() {
		op := p.parseOperation()
		loop.Operations = append(loop.Operations, op)
	}

	if p.tok == token.EOF {
		p.error(pos, `unexpected EOF, expected ]`)
	}

	return loop
}

func (p *Parser) error(pos token.Position, err string) {
	p.ErrorCount++
	if p.err != nil {
		p.err(pos, err)
	}
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.l.Next()
}
