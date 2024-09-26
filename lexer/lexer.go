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
	case ',':
		tok = newToken(token.COMMA, string(lex.ch))
	case ':':
		tok = newToken(token.COLON, string(lex.ch))
	case ';':
		tok = newToken(token.SEMICOLON, string(lex.ch))
	case '(':
		tok = newToken(token.LPAREN, string(lex.ch))
	case ')':
		tok = newToken(token.RPAREN, string(lex.ch))
	case '{':
		tok = newToken(token.LBRACE, string(lex.ch))
	case '}':
		tok = newToken(token.RBRACE, string(lex.ch))
	case '"':
		tok = newToken(token.STRING, lex.readString())
	case '-':
		if lex.peekChar() == '-' {
			ch := lex.ch
			lex.readChar()
			tok = newToken(token.CUSTOM_NAME_PREFIX, string(ch)+string(lex.ch))
		}
	case 0:
		tok = newToken(token.EOF, "")
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		tok = newToken(token.ILLEGAL, string(lex.ch))
	}

	lex.readChar()

	return tok
}
