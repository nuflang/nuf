package lexer_test

import (
	"testing"

	"github.com/nuflang/nuf/lexer"
)

func TestTokenize(t *testing.T) {
	source := "> banner"
	tokens := lexer.Tokenize(source)
	if len(tokens) != 3 {
		t.Errorf("Expected 3 tokens, Got: %d tokens", len(tokens))
	}

	if tokens[0].Kind != 1 {
		t.Errorf("Expected: 1, Got: %d", tokens[0].Kind)
	}
	if tokens[0].Value != ">" {
		t.Errorf("Expected: '>', Got: %s", tokens[0].Value)
	}

	if tokens[1].Kind != 2 {
		t.Errorf("Expected: 2, Got: %d", tokens[1].Kind)
	}
	if tokens[1].Value != "banner" {
		t.Errorf("Expected: 'banner', Got: %s", tokens[1].Value)
	}

	if tokens[2].Kind != 0 {
		t.Errorf("Expected: 0, Got: %d", tokens[2].Kind)
	}
	if tokens[2].Value != "EOF" {
		t.Errorf("Expected: 'EOF', Got: %s", tokens[2].Value)
	}
}
