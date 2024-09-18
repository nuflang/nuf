package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = iota
	EOF

	SEMICOLON

	STRING
)
