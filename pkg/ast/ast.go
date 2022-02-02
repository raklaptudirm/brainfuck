package ast

import "github.com/raklaptudirm/brainfuck/pkg/token"

type Node interface {
	TokenLiteral() string
}

type Program struct {
	Operations []Operation
}

func (p *Program) TokenLiteral() string {
	return "program"
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
	Operators []Operation
}

func (l *Loop) TokenLiteral() string {
	return "loop"
}

func (l *Loop) operationNode() {}
