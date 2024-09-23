package lexer

import (
	"testing"

	"github.com/nuflang/nuf/token"
)

func TestStringToken(t *testing.T) {
	input := `"Paragraph";`
	lex := NewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "Paragraph"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d. Got %d.", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s. Got %s.", test.expectedLiteral, tok.Literal)
		}
	}
}

func TestFunctionCallTokens(t *testing.T) {
	input := `section_title("Heading");`
	lex := NewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "section_title"},
		{token.LPAREN, "("},
		{token.STRING, "Heading"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d. Got %d.", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s. Got %s.", test.expectedLiteral, tok.Literal)
		}
	}
}

func TestKeywords(t *testing.T) {
	input := `inside`
	lex := NewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INSIDE, "inside"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d. Got %d.", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s. Got %s.", test.expectedLiteral, tok.Literal)
		}
	}
}

func TestCustomName(t *testing.T) {
	input := `--custom_name`
	lex := NewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CUSTOM_NAME_PREFIX, "--"},
		{token.IDENT, "custom_name"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d. Got %d.", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s. Got %s.", test.expectedLiteral, tok.Literal)
		}
	}
}

func TestElementInsideElement(t *testing.T) {
	input := `section_title("Heading 1") inside --main;`
	lex := NewLexer(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "section_title"},
		{token.LPAREN, "("},
		{token.STRING, "Heading 1"},
		{token.RPAREN, ")"},
		{token.INSIDE, "inside"},
		{token.CUSTOM_NAME_PREFIX, "--"},
		{token.IDENT, "main"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d. Got %d.", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s. Got %s.", test.expectedLiteral, tok.Literal)
		}
	}
}
