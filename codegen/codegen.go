package codegen

import (
	"strconv"

	"github.com/nuflang/nuf/ast"
)

func getHTMLOpenTag(tag string) string {
	return "<" + tag + ">"
}

func getHTMLCloseTag(tag string) string {
	return "</" + tag + ">"
}

func GenerateHTML(program ast.Program) string {
	html := ""

	for _, block := range program.Page.Body {
		html += getHTMLOpenTag(block.(ast.LandmarkStatement).HTMLTag)

		for _, heading := range block.(ast.LandmarkStatement).Body {
			html += "\n    "
			html += generateHeadingHTML(heading.(ast.HeadingStatement))
		}

		html += "\n"
		html += getHTMLCloseTag(block.(ast.LandmarkStatement).HTMLTag)
		html += "\n"
	}

	return html
}

func generateHeadingHTML(heading ast.HeadingStatement) string {
	html := ""

	html += getHTMLOpenTag("h" + strconv.Itoa(heading.HeadingLevel))
	html += getHTMLCloseTag("h" + strconv.Itoa(heading.HeadingLevel))

	return html
}
