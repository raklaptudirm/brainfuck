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
