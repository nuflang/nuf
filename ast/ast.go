package ast

type Program struct {
	Page Page
}

type Page struct {
	Body []BlockStatement
}

type BlockStatement interface {
	blockStatement()
}

type LandmarkStatement struct {
	HTMLTag string
	Body    []BlockStatement
}

func (node LandmarkStatement) blockStatement() {}

type HeadingStatement struct {
	HeadingLevel int
}

func (node HeadingStatement) blockStatement() {}
