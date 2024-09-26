package parser

import (
	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
	"github.com/nuflang/nuf/token"
)

const (
	_ byte = iota
	LOWEST
	INSIDE
	CALL // function()
)

var precedences = map[token.TokenType]byte{
	token.INSIDE: INSIDE,
	token.LPAREN: CALL,
}

type Parser struct {
	lex            *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
	errors         []string
}

func NewParser(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    lex,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.CUSTOM_NAME_PREFIX, p.parseCustomNameExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.INSIDE, p.parseInfixExpression)

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

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		Token:      p.currentToken,
		Expression: p.parseExpression(LOWEST),
	}

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) parseExpression(precedence byte) ast.Expression {
	prefixFn := p.prefixParseFns[p.currentToken.Type]
	if prefixFn == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}

	leftExpression := prefixFn()

	for p.peekToken.Type != token.SEMICOLON && precedence < p.peekPrecedence() {
		infixFn := p.infixParseFns[p.peekToken.Type]
		if infixFn == nil {
			return leftExpression
		}

		p.nextToken()
		leftExpression = infixFn(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) parseCallExpression(fn ast.Expression) ast.Expression {
	return &ast.CallExpression{
		Token:     p.currentToken,
		Function:  fn,
		Arguments: p.parseCallArguments(),
	}
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekToken.Type == token.RPAREN {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseCustomNameExpression() ast.Expression {
	p.nextToken()

	expression := &ast.CustomNameExpression{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.nextToken()

	expression.Right = p.parseExpression(precedence)

	return expression
}
