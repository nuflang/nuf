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

	if output.Value != "<p>Paragraph</p>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<p>Paragraph</p>", output.Value)
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

	if output.Value != "<h1>Heading</h1>" {
		t.Fatalf("Expected `%s`. Got `%s`.", "<h1>Heading</h1>", output.Value)
	}
}
