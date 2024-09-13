package codegen

import (
	"github.com/nuflang/nuf/ast"
)

func getHTMLOpenTag(tag string) string {
	return "<" + tag + ">"
}

func getHTMLCloseTag(tag string) string {
	return "</" + tag + ">"
}

func GenerateHTML(ast ast.BlockStatement) string {
	html := ""

	for _, block := range ast.Body {
		html += getHTMLOpenTag(block.HTMLTag)
		html += getHTMLCloseTag(block.HTMLTag)
		html += "\n"
	}

	return html
}
