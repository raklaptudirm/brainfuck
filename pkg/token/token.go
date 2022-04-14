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

// Package token defines constants which represent the basic lexical
// elements of the brainfuck programming language. It also defines types
// which represent the context a token was found in.
package token

import "fmt"

// Token represents the basic lexical element of the brainfuck programming
// language and the context it was found in.
type Token struct {
	Type
	Position
}

// Type represents the type of a Token.
type Type int

const (
	Eof Type = iota

	Plus  // +
	Minus // -

	LeftArrow  // <
	RightArrow // >

	Comma  // ,
	Period // .

	LeftBracket  // [
	RightBracket // ]
)

var tokens = [...]string{
	Eof:          "EOF",
	Plus:         "+",
	Minus:        "-",
	LeftArrow:    "<",
	RightArrow:   ">",
	Comma:        ",",
	Period:       ".",
	LeftBracket:  "[",
	RightBracket: "]",
}

// String returns a string representation of the Type.
func (t Type) String() string {
	return tokens[t]
}

// Position represents the position of a token in a file.
type Position struct {
	Line   int
	Column int
}

// String returns a string representation of a position in the format
// <line>:<column>
func (p Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

// NextLine moves the position to the next line.
func (p *Position) NextLine() {
	p.Line++
	p.Column = 1
}
