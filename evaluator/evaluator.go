package evaluator

import (
	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/object"
)

type Output struct {
	Value string
}

func NewOutput() *Output {
	return &Output{
		Value: "",
	}
}

func (o *Output) Eval(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		o.evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		o.Eval(node.Expression)
	case *ast.StringLiteral:
		s := object.String{Value: node.Value}
		o.Value += "<p>" + s.Inspect() + "</p>\n"
	}
}

func (o *Output) evalStatements(statements []ast.Statement) {
	for _, statement := range statements {
		o.Eval(statement)
	}
}
