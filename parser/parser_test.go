package parser

import (
	"testing"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `"Hello, World!"`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	statement := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := statement.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("Wrong expression type. Expected *ast.StringLiteral. Got %T.", statement.Expression)
	}
	if literal.Value != "Hello, World!" {
		t.Errorf("Wrong literal value. Expected %s. Got %q.", "Hello, World!", literal.Value)
	}
}
