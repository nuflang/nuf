package parser

import (
	"testing"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
)

func TestStringLiteralExpression(t *testing.T) {
	input := `"Paragraph";`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	statement := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := statement.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("Expected `StringLiteral`. Got `%T`.", statement.Expression)
	}
	if literal.Value != "Paragraph" {
		t.Errorf("Expected `%s`. Got `%q`.", "Paragraph", literal.Value)
	}
}

func TestCallExpression(t *testing.T) {
	input := `section_title("Heading");`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("Program.Statements does not contain `%d` statement. Got `%d`.", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected `ExpressionStatement` statement. Got `%T`", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("Expected `CallExpression` expression. Got `%T`.", statement.Expression)
	}

	if !testIdentifier(t, expression.Function, "section_title") {
		return
	}

	if len(expression.Arguments) != 1 {
		t.Fatalf("Expected `%d` argument. Got `%d`.", 1, len(expression.Arguments))
	}

	literal := expression.Arguments[0]
	if !ok {
		t.Fatalf("Expected `StringLiteral`. Got `%T`.", statement.Expression)
	}
	if literal.TokenLiteral() != "Heading" {
		t.Errorf("Expected `%s`. Got `%q`.", "Heading", literal.TokenLiteral())
	}
}

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	identifier, ok := expression.(*ast.Identifier)

	if !ok {
		t.Errorf("Expected `Identifier` expression. Got %v, %v, `%T`.", identifier, ok, expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("Expected identifier value `%s`. Got `%s`.", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("Expected identifier token literal `%s`. Got `%s`.", value, identifier.TokenLiteral())
		return false
	}

	return true
}
