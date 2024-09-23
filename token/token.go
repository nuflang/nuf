package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = iota
	EOF

	IDENT
	STRING

	SEMICOLON

	LPAREN
	RPAREN

	CUSTOM_NAME_PREFIX // --

	// Keywords
	INSIDE
)

var keywords = map[string]TokenType{
	"inside": INSIDE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
