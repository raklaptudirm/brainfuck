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

// Package ast defines structures to represent brainfuck code as an
// abstract syntax tree.
package ast

import "laptudirm.com/x/brainfuck/pkg/token"

// Node is the interface implemented by all ast nodes.
type Node interface {
	Node()
}

// Program is the main node of a brainfuck ast.
type Program struct {
	Operations []Operation
}

func (p *Program) Node() {}

// Operation is the interface implemented by all brainfuck operation ast
// nodes.
type Operation interface {
	Node
	Operation()
}

// Loop is the ast node which represents a brainfuck loop.
type Loop struct {
	Operations []Operation
}

func (l *Loop) Node()      {}
func (l *Loop) Operation() {}

// Operator is the ast node which represents all the simple operations.
type Operator token.Token

func (o *Operator) Node()      {}
func (o *Operator) Operation() {}
