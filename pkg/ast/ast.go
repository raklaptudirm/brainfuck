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

type Operator token.Token

func (o *Operator) TokenLiteral() string {
	return token.Token(*o).String()
}

func (o *Operator) operationNode() {}

type Loop struct {
	Operators []Operation
}

func (l *Loop) TokenLiteral() string {
	return "loop"
}

func (l *Loop) operationNode() {}
