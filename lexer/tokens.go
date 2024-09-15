package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota

	OPEN_CURLY  // {
	CLOSE_CURLY // }

	GREATER // Landmark
	PLUS    // Heading

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

	// Headings
	H1
	H2
	H3
	H4
	H5
	H6
)

var reservedKeywordLookup map[string]TokenKind = map[string]TokenKind{
	// Landmarks
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

	// Headings
	"h1": H1,
	"h2": H2,
	"h3": H3,
	"h4": H4,
	"h5": H5,
	"h6": H6,
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
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
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
	case H1:
		return "h1"
	case H2:
		return "h2"
	case H3:
		return "h3"
	case H4:
		return "h4"
	case H5:
		return "h5"
	case H6:
		return "h6"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
