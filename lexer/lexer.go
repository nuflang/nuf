package lexer

import "github.com/nuflang/nuf/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte //current char under examination
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()

	return lex
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case ';':
		tok = newToken(token.SEMICOLON, string(lex.ch))
	case '(':
		tok = newToken(token.LPAREN, string(lex.ch))
	case ')':
		tok = newToken(token.RPAREN, string(lex.ch))
	case '"':
		tok = newToken(token.STRING, lex.readString())
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(lex.ch) {
			tok = newToken(token.LookupIdent(tok.Literal), lex.readIdentifier())
			return tok
		}

		tok = newToken(token.ILLEGAL, string(lex.ch))
	}

	lex.readChar()

	return tok
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) readString() string {
	position := lex.position + 1

	for {
		lex.readChar()

		if lex.ch == '"' || lex.ch == 0 {
			break
		}
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'
}
