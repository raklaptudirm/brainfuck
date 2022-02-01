package token

import "strconv"

type Token int

const (
	// special tokens
	ILLEGAL Token = iota
	COMMENT
	EOF

	// value operators
	INC_VAL // +
	DEC_VAL // -

	// memory address
	INC_PTR // >
	DEC_PTR // <

	// looping
	SLOOP // [
	ELOOP // ]

	// io operators
	INPUT // ,
	PRINT // .
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	COMMENT: "COMMENT",
	EOF:     "EOF",

	INC_VAL: "+",
	DEC_VAL: "-",

	INC_PTR: ">",
	DEC_PTR: "<",

	SLOOP: "[",
	ELOOP: "]",

	INPUT: ",",
	PRINT: ".",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}
