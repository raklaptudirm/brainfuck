package lexer

import (
	"unicode/utf8"

	"github.com/raklaptudirm/brainfuck/pkg/token"
)

type lexer struct {
	src string
	ch  rune

	offset   int
	rdOffset int

	line int
	col  int
}

const eof = 0

func (l *lexer) Next() (pos int, tok token.Token, lit string) {
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

func (l *lexer) lexComment() token.Token {
	for ch := l.peek(); !isOperator(ch) && !l.atEnd(); ch = l.peek() {
		l.consume()
	}

	return token.COMMENT
}

func (l *lexer) consume() {
	if l.ch == '\n' {
		l.line++
		l.col = 0
	}

	if l.atEnd() {
		l.ch = eof
		return
	}

	r, w := rune(l.src[l.rdOffset]), 1
	if r >= utf8.RuneSelf {
		r, w = utf8.DecodeRuneInString(l.src[l.rdOffset:])
	}

	l.ch = r

	l.rdOffset += w
	l.col += w
}

func (l *lexer) peek() byte {
	if l.atEnd() {
		return eof
	}

	return l.src[l.rdOffset]
}

func (l *lexer) atEnd() bool {
	return l.rdOffset >= len(l.src)
}

func New(s string) *lexer {
	return &lexer{
		src: s,

		offset:   0,
		rdOffset: 0,

		line: 0,
		col:  0,
	}
}

func isOperator(b byte) bool {
	switch b {
	case '+', '-', '>', '<', '[', ']', ',', '.':
		return true
	default:
		return false
	}
}
