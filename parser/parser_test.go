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

func TestCallExpressionWithMultipleArguments(t *testing.T) {
	input := `section_title("Heading", "Argument");`

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

	if len(expression.Arguments) != 2 {
		t.Fatalf("Expected `%d` argument. Got `%d`.", 2, len(expression.Arguments))
	}

	literal := expression.Arguments[0]
	if !ok {
		t.Fatalf("Expected `StringLiteral`. Got `%T`.", statement.Expression)
	}
	if literal.TokenLiteral() != "Heading" {
		t.Errorf("Expected `%s`. Got `%q`.", "Heading", literal.TokenLiteral())
	}

	literal = expression.Arguments[1]
	if !ok {
		t.Fatalf("Expected `StringLiteral`. Got `%T`.", statement.Expression)
	}
	if literal.TokenLiteral() != "Argument" {
		t.Errorf("Expected `%s`. Got `%q`.", "Argument", literal.TokenLiteral())
	}
}

func TestCallExpressionWithOptions(t *testing.T) {
	input := `section("region", { name: --custom_region, key: "value" });`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected `ExpressionStatement` statement. Got `%T`", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("Expected `CallExpression` expression. Got `%T`.", statement.Expression)
	}

	if !testIdentifier(t, expression.Function, "section") {
		return
	}

	if len(expression.Arguments) != 2 {
		t.Fatalf("Expected `%d` argument. Got `%d`.", 2, len(expression.Arguments))
	}

	literal := expression.Arguments[0]
	if !ok {
		t.Fatalf("Expected `Identifier`. Got `%T`.", statement.Expression)
	}
	if literal.TokenLiteral() != "region" {
		t.Errorf("Expected `%s`. Got `%q`.", "region", literal.TokenLiteral())
	}

	hashExpression := expression.Arguments[1]
	hash, ok := hashExpression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("Expected `HashLiteral`. Got `%T`.", hashExpression)
	}

	if len(hash.Pairs) != 2 {
		t.Errorf("Expected hash length to be `%d`. Got `%d`.", 2, len(hash.Pairs))
	}

	expectedHash := map[string]func(ast.Expression){
		"name": func(e ast.Expression) {
			customNameExpression, ok := e.(*ast.CustomNameExpression)
			if !ok {
				t.Fatalf("Expected `CustomNameExpression` expression. Got `%T`.", statement.Expression)
			}

			if customNameExpression.Value != "custom_region" {
				t.Fatalf("Expected value `%s`. Got `%s`.", "custom_region", customNameExpression.Value)
			}
		},
		"key": func(e ast.Expression) {
			stringLiteralExpression, ok := e.(*ast.StringLiteral)
			if !ok {
				t.Fatalf("Expected `StringLiteral` expression. Got `%T`.", statement.Expression)
			}

			if stringLiteralExpression.Value != "value" {
				t.Fatalf("Expected value `%s`. Got `%s`.", "value", stringLiteralExpression.Value)
			}
		},
	}

	for key, value := range hash.Pairs {
		identifier, ok := key.(*ast.Identifier)
		if !ok {
			t.Errorf("Key is not `Identifier`. Got `%T`.", key)
		}

		testFn, ok := expectedHash[identifier.Value]
		if !ok {
			t.Errorf("No test function for key `%q` found", identifier.Value)
			continue
		}

		testFn(value)
	}
}

func TestInfixExpression(t *testing.T) {
	input := `section_title("Heading 1") inside --main;`

	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("Program.Statements does not contain `%d` statement. Got `%d`.", 1, len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Expected `CallExpression` statement. Got `%T`", program.Statements[0])
	}

	expression, ok := statement.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Expected `InfixExpression` expression. Got `%T`.", statement.Expression)
	}

	if expression.Operator != "inside" {
		t.Fatalf("Expected `%s` operator. Got `%s`.", "inside", expression.Operator)
	}

	customNameExpression, ok := expression.Right.(*ast.CustomNameExpression)
	if !ok {
		t.Fatalf("Expected `CustomNameExpression` expression. Got `%T`.", statement.Expression)
	}

	if customNameExpression.Value != "main" {
		t.Fatalf("Expected value `%s`. Got `%s`.", "main", customNameExpression.Value)
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
