package lexer

import (
	"fmt"
	"unicode/utf8"

	"github.com/raklaptudirm/brainfuck/pkg/token"
)

type Lexer struct {
	src string
	ch  rune

	err ErrorHandler

	offset   int
	rdOffset int

	line int
	col  int
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

	l.line = 0
	l.col = 0
}

func (l *Lexer) Next() (pos int, tok token.Token, lit string) {
	switch l.peek() {
	case eof:
		l.consume()
		tok = token.EOF
	case '+':
		l.consume()
		tok = token.INC_VAL
	case '-':
		l.consume()
		tok = token.DEC_VAL
	case '>':
		l.consume()
		tok = token.INC_PTR
	case '<':
		l.consume()
		tok = token.DEC_PTR
	case ',':
		l.consume()
		tok = token.INPUT
	case '.':
		l.consume()
		tok = token.PRINT
	case '[':
		l.consume()
		tok = token.SLOOP
	case ']':
		l.consume()
		tok = token.ELOOP
	default:
		tok = l.lexComment()
	}

	pos = l.offset
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
	if l.ch == '\n' {
		l.line++
		l.col = 0
	}

	if l.atEnd() {
		l.ch = eof
		return
	}

	r, w := rune(l.src[l.rdOffset]), 1
	if r == 0 {
		l.error("illegal character NUL")
		goto advance
	}

	if r >= utf8.RuneSelf {
		r, w = utf8.DecodeRuneInString(l.src[l.rdOffset:])

		if r == utf8.RuneError && w == 1 {
			l.error("illegal UTF-8 encoding")
			goto advance
		}

		if r == bom && l.offset > 0 {
			l.error("illegal byte order mark")
			goto advance
		}
	}

advance:
	l.ch = r

	l.rdOffset += w
	l.col += w
}

func (l *Lexer) error(err string) {
	if l.err != nil {
		l.err(l.line, l.col, err)
	}
}

func (l *Lexer) errorf(format string, a ...interface{}) {
	err := fmt.Sprintf(format, a...)
	l.error(err)
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

type ErrorHandler func(int, int, string)

func isOperator(r rune) bool {
	switch r {
	case '+', '-', '>', '<', '[', ']', ',', '.':
		return true
	default:
		return false
	}
}
