package lexer

import (
	"testing"

	"github.com/nuflang/nuf/token"
)

func TestStringToken(t *testing.T) {
	input := `  "Hello, World!"  `
	lex := NewLexer(input)

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.STRING, "Hello, World!"},
		{token.EOF, ""},
	}

	for _, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("Wrong token type. Expected %d, got %d", test.expectedType, tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("Wrong token literal. Expected %s, got %s", test.expectedLiteral, tok.Literal)
		}
	}
}
