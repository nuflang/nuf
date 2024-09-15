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

func getHeadingLevel(kind lexer.TokenKind) int {
	switch kind {
	case lexer.H1:
		return 1
	case lexer.H2:
		return 2
	case lexer.H3:
		return 3
	case lexer.H4:
		return 4
	case lexer.H5:
		return 5
	case lexer.H6:
		return 6
	default:
		return 0
	}
}

func parseStatement(p *parser) ast.BlockStatement {
	statementFunc, exists := blockStatementLookupTable[p.currentTokenKind()]

	if !exists {
		fmt.Printf("Failed to parse statement: %d", p.currentTokenKind())
		return nil
	}

	return statementFunc(p)
}

func parseLandmarkBlockStatement(p *parser) ast.BlockStatement {
	p.expect(lexer.GREATER)
	tokenKind := p.advance().Kind
	landmarkRole := getLandmarkHTMLTag(tokenKind)

	body := make([]ast.BlockStatement, 0)

	p.expect(lexer.OPEN_CURLY)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		body = append(body, parseStatement(p))
	}

	p.expect(lexer.CLOSE_CURLY)

	return ast.LandmarkStatement{
		HTMLTag: landmarkRole,
		Body:    body,
	}
}

func parseHeadingBlockStatement(p *parser) ast.BlockStatement {
	p.expect(lexer.PLUS)
	tokenKind := p.advance().Kind
	headingLevel := getHeadingLevel(tokenKind)

	return ast.HeadingStatement{
		HeadingLevel: headingLevel,
	}
}
