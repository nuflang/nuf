package parser

import (
	"fmt"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/lexer"
)

func getLandmarkHTMLTag(kind lexer.TokenKind) string {
	switch kind {
	case lexer.BANNER:
		return "header"
	case lexer.HEADER:
		return "header"
	case lexer.COMPLEMENTARY:
		return "aside"
	case lexer.ASIDE:
		return "aside"
	case lexer.CONTENTINFO:
		return "footer"
	case lexer.FOOTER:
		return "footer"
	case lexer.FORM:
		return "form"
	case lexer.MAIN:
		return "main"
	case lexer.NAVIGATION:
		return "nav"
	case lexer.NAV:
		return "nav"
	case lexer.REGION:
		return "section"
	case lexer.SECTION:
		return "section"
	case lexer.SEARCH:
		return "search"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

func parseLandmarkBlockStatement(p *parser) ast.BlockStatement {
	p.expect(lexer.GREATER)
	tokenKind := p.advance().Kind
	landmarkRole := getLandmarkHTMLTag(tokenKind)

	return ast.BlockStatement{
		HTMLTag: landmarkRole,
	}
}
