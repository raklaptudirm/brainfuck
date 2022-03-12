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

package ast

import "github.com/raklaptudirm/brainfuck/pkg/token"

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Operations []Operation
}

func (p *Program) TokenLiteral() string {
	s := ""
	for _, op := range p.Operations {
		s += op.TokenLiteral()
	}

	return s
}

type Operation interface {
	Node
	operationNode()
}

type Comment struct {
	Literal string
}

func (c *Comment) TokenLiteral() string {
	return c.Literal
}

func (c *Comment) operationNode() {}

type Operator struct {
	Token token.Token
}

func (o *Operator) TokenLiteral() string {
	return o.Token.String()
}

func (o *Operator) operationNode() {}

type Loop struct {
	Operations []Operation
}

func (l *Loop) TokenLiteral() string {
	s := "["
	for _, op := range l.Operations {
		s += op.TokenLiteral()
	}

	return s + "]"
}

func (l *Loop) operationNode() {}
