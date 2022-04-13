package lexer_test

import (
	"testing"

	"laptudirm.com/x/brainfuck/pkg/lexer"
	"laptudirm.com/x/brainfuck/pkg/token"
)

var test = `Hello World
+- foo
<> bar
,. baz
[] etc
`
var output = []token.Token{
	{Type: token.Plus, Position: token.Position{Line: 2, Column: 1}},
	{Type: token.Minus, Position: token.Position{Line: 2, Column: 2}},
	{Type: token.LeftArrow, Position: token.Position{Line: 3, Column: 1}},
	{Type: token.RightArrow, Position: token.Position{Line: 3, Column: 2}},
	{Type: token.Comma, Position: token.Position{Line: 4, Column: 1}},
	{Type: token.Period, Position: token.Position{Line: 4, Column: 2}},
	{Type: token.LeftBracket, Position: token.Position{Line: 5, Column: 1}},
	{Type: token.RightBracket, Position: token.Position{Line: 5, Column: 2}},
	{Type: token.Eof, Position: token.Position{Line: 6, Column: 1}},
}

func TestLexer(t *testing.T) {
	ch := lexer.Lex([]byte(test))
	i := 0
	for tok := range ch {
		exp := output[i]

		if tok.Type != exp.Type {
			t.Fatalf("token %d: expected %s, received %s\n", i, exp.Type, tok.Type)
		}

		if tok.Position != exp.Position {
			t.Fatalf("token %d: expected %s, received %s", i, exp.Position, tok.Position)
		}

		i++
	}
}
