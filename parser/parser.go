package parser

import (
	"fmt"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
	// errors []error
}

func Parse(tokens []lexer.Token) ast.Program {
	program := ast.Program{Page: ast.Page{Body: make([]ast.BlockStatement, 0)}}
	parser := createParser(tokens)

	for parser.hasTokens() {
		program.Page.Body = append(program.Page.Body, parseLandmarkBlockStatement(parser))
	}

	return program
}

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()

	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s, but received %s", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
