package evaluator

import (
	"testing"

	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/object"
	"github.com/nuflang/nuf/parser"
)

func TestParagraph(t *testing.T) {
	input := `"Paragraph"`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	program := p.ParseProgram()
	output := NewOutput()
	env := object.NewEnvironment()
	output.Eval(program, env)

	if output.HTMLValue != "<p>Paragraph</p>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<p>Paragraph</p>", output.HTMLValue)
	}
}

func TestSectionTitleBuiltinFunction(t *testing.T) {
	input := `section_title("Heading");`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	program := p.ParseProgram()
	output := NewOutput()
	env := object.NewEnvironment()
	output.Eval(program, env)

	if output.HTMLValue != "<h1>Heading</h1>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<h1>Heading</h1>", output.HTMLValue)
	}
}

func TestSectionBuiltinFunction(t *testing.T) {
	input := `section("main");`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	program := p.ParseProgram()
	output := NewOutput()
	env := object.NewEnvironment()
	output.Eval(program, env)

	if output.HTMLValue != "<main></main>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<main></main>", output.HTMLValue)
	}
}

func TestInfixExpressionInside(t *testing.T) {
	input := `section_title("Heading 1") inside --main;section("main")`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	program := p.ParseProgram()
	output := NewOutput()
	env := object.NewEnvironment()
	output.Eval(program, env)

	if output.HTMLValue != "<main><h1>Heading</h1></main>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<main><h1>Heading</h1></main>", output.HTMLValue)
	}
}
