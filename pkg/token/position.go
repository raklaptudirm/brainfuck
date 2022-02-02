package token

import "fmt"

type Position struct {
	Line int
	Col  int
}

func (p *Position) String() string {
	return fmt.Sprintf("%v:%v", p.Line, p.Col)
}

func (p *Position) NextLine() {
	p.Line++
	p.Col = 1
}
