package parser

import (
	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/token"
)

type prefixParseFn func() ast.Expression

type Parser struct {
	lex            *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
}

func NewParser(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:            lex,
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
	}

	p.registerPrefix(token.STRING, p.parseStringLiteral)

	// Read two tokens, so currentToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currentToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: p.parseExpression(),
	}

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression() ast.Expression {
	prefixFn := p.prefixParseFns[p.currentToken.Type]

	if prefixFn == nil {
		return nil
	}

	return prefixFn()
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}
