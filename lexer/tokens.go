package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota

	GREATER // Landmark

	// Reserved keywords

	// Landmarks
	BANNER
	HEADER
	COMPLEMENTARY
	ASIDE
	CONTENTINFO
	FOOTER
	FORM
	MAIN
	NAVIGATION
	NAV
	REGION
	SECTION
	SEARCH
)

var reservedKeywordLookup map[string]TokenKind = map[string]TokenKind{
	"banner":        BANNER,
	"header":        HEADER,
	"complementary": COMPLEMENTARY,
	"aside":         ASIDE,
	"contentinfo":   CONTENTINFO,
	"footer":        FOOTER,
	"form":          FORM,
	"main":          MAIN,
	"navigation":    NAVIGATION,
	"nav":           NAV,
	"region":        REGION,
	"section":       SECTION,
	"search":        SEARCH,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) Debug() {
	fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case GREATER:
		return "greater"
	case BANNER:
		return "banner"
	case HEADER:
		return "header"
	case COMPLEMENTARY:
		return "complementary"
	case ASIDE:
		return "aside"
	case CONTENTINFO:
		return "contentinfo"
	case FOOTER:
		return "footer"
	case FORM:
		return "form"
	case MAIN:
		return "main"
	case NAVIGATION:
		return "navigation"
	case NAV:
		return "nav"
	case REGION:
		return "region"
	case SECTION:
		return "section"
	case SEARCH:
		return "search"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
