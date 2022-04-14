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

// Package lexer contains an implementation of a brainfuck lexer which
// lexes brainfuck code into tokens concurrently. It only exposes the Lex
// function which should be used to lex any brainfuck code.
package lexer

import "laptudirm.com/x/brainfuck/pkg/token"

// Lex lexes brainfuck code into a stream of tokens concurrently which are
// sent into the tokens channel. It does not verify whether the code is
// valid brainfuck.
func Lex(data []byte) <-chan token.Token {
	l := lexer{
		data:   data,
		pos:    token.Position{Line: 1, Column: 1},
		tokens: make(chan token.Token),
	}

	// start concurrent tokenization
	go l.program()
	return l.tokens
}

// lexer is a state machine representing the current state of the lexer.
type lexer struct {
	data []byte // source data

	// state
	offset int              // current offset within data
	curr   rune             // current rune
	pos    token.Position   // current position
	tokens chan token.Token // tokens channel
}

// eof is a constant representing the end of file.
const eof = -1

// program lexes a brainfuck program from the data in the lexer.
func (l *lexer) program() {
	for {
		var tok token.Type

		switch l.peek() {
		case eof:
			tok = token.Eof
		case '+':
			tok = token.Plus
		case '-':
			tok = token.Minus
		case '<':
			tok = token.LeftArrow
		case '>':
			tok = token.RightArrow
		case ',':
			tok = token.Comma
		case '.':
			tok = token.Period
		case '[':
			tok = token.LeftBracket
		case ']':
			tok = token.RightBracket
		default:
			// ignore comments
			l.next()
			continue
		}

		// emit token
		l.emit(tok)
		l.next()
		if tok == token.Eof {
			return
		}
	}
}

// next moves the lexer forward in it's data.
func (l *lexer) next() {
	if l.curr = l.peek(); l.curr == eof {
		return
	}

	l.offset++

	if l.curr == '\n' {
		l.pos.NextLine()
	} else {
		l.pos.Column++
	}
}

// peek returns the next byte in the lexer's data.
func (l *lexer) peek() rune {
	// check if at eof
	if l.offset >= len(l.data) {
		return eof
	}

	return rune(l.data[l.offset])
}

// emit emits the a token of the provided token type into the token stream.
func (l *lexer) emit(tok token.Type) {
	l.tokens <- token.Token{
		Type:     tok,
		Position: l.pos,
	}

	// if Eof has been emitted, close channel
	if tok == token.Eof {
		close(l.tokens)
	}
}
