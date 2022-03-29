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

package lexer

import (
	"unicode/utf8"

	"laptudirm.com/x/brainfuck/pkg/token"
)

type Lexer struct {
	src string
	ch  rune

	err ErrorHandler

	offset   int
	rdOffset int

	pos token.Position

	ErrorCount int
}

const (
	eof = -1     // end of file
	bom = 0xFEFF // byte order mark
)

func (l *Lexer) Init(src string, handler ErrorHandler) {
	l.src = src
	l.err = handler

	l.offset = 0
	l.rdOffset = 0

	l.pos = token.Position{
		Line: 1,
		Col:  1,
	}
}

func (l *Lexer) Next() (pos token.Position, tok token.Token, lit string) {
	pos = l.pos

	switch l.consume(); l.ch {
	case eof:
		tok = token.EOF
	case '+':
		tok = token.INC_VAL
	case '-':
		tok = token.DEC_VAL
	case '>':
		tok = token.INC_PTR
	case '<':
		tok = token.DEC_PTR
	case ',':
		tok = token.INPUT
	case '.':
		tok = token.PRINT
	case '[':
		tok = token.SLOOP
	case ']':
		tok = token.ELOOP
	default:
		tok = l.lexComment()
	}

	lit = l.src[l.offset:l.rdOffset]
	l.offset = l.rdOffset
	return
}

func (l *Lexer) lexComment() token.Token {
	for ch := l.peek(); !isOperator(ch) && !l.atEnd(); ch = l.peek() {
		l.consume()
	}

	return token.COMMENT
}

func (l *Lexer) consume() {
	if l.atEnd() {
		l.ch = eof
		return
	}

	r, w := rune(l.src[l.rdOffset]), 1
	if r == 0 {
		l.error("illegal character NUL")
		goto advance
	}

	if r < utf8.RuneSelf {
		goto advance
	}

	r, w = utf8.DecodeRuneInString(l.src[l.rdOffset:])

	if r == utf8.RuneError && w == 1 {
		l.error("illegal UTF-8 encoding")
		goto advance
	}

	if r == bom && l.offset > 0 {
		l.error("illegal byte order mark")
	}

advance:
	l.ch = r

	l.rdOffset += w
	l.pos.Col += w

	if r == '\n' {
		l.pos.NextLine()
	}
}

func (l *Lexer) error(err string) {
	l.ErrorCount++
	if l.err != nil {
		l.err(l.pos, err)
	}
}

func (l *Lexer) peek() rune {
	if l.atEnd() {
		return eof
	}

	return rune(l.src[l.rdOffset])
}

func (l *Lexer) atEnd() bool {
	return l.rdOffset >= len(l.src)
}

type ErrorHandler func(token.Position, string)

func isOperator(r rune) bool {
	switch r {
	case '+', '-', '>', '<', '[', ']', ',', '.':
		return true
	default:
		return false
	}
}
