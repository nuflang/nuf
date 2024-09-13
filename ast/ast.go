package ast

type BlockStatement struct {
	HTMLTag string
	Body    []BlockStatement
}
